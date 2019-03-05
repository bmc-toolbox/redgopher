package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/bmc-toolbox/redgopher/openapi"
)

func main() {

	var user, pass string

	user = "foo"
	pass = "bar"
	ip := "127.0.0.1"

	//Since BMCs have shitty certs
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// declare behaviour on redirects
	checkRedirect := func(req *http.Request, via []*http.Request) error {
		if len(via) >= 5 {
			log.Printf("Too many redirects - %d\n", len(via))
		}
		// Sensitive headers are not forwarded by default
		req.SetBasicAuth(user, pass)
		return nil
	}

	cfg := &openapi.Configuration{
		BasePath:   fmt.Sprintf("https://%s", ip),
		HTTPClient: &http.Client{Transport: tr, CheckRedirect: checkRedirect},
		UserAgent:  "redgopher",
	}

	client := openapi.NewAPIClient(cfg)

	ctx := context.Background()
	basicAuth := openapi.BasicAuth{UserName: user, Password: pass}
	ctx2 := context.WithValue(ctx, openapi.ContextBasicAuth, basicAuth)

	//s, _, err := client.DefaultApi.RedfishV1SessionServiceSessionsGet(ctx2)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//log.Printf("%+v", s)

	l, _, err := client.DefaultApi.RedfishV1ManagersManagerIdNetworkProtocolGet(ctx2, "1")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", l)
}
