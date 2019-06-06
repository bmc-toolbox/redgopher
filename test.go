package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	ctx = context.WithValue(ctx, openapi.ContextBasicAuth, basicAuth)

	// here we get the manager we will need to use later
	m, _, err := client.DefaultApi.RedfishV1ManagersGet(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response: %+v", m)

	// for the manager we will use only the last part of the id
	manager := strings.Split(m.Members[0].OdataId, "/")[strings.Count(m.Members[0].OdataId, "/")]
	fmt.Printf("manager: %s\n", manager)

	// here we use the manager as argument to retrieve the required data
	l, _, err := client.DefaultApi.RedfishV1ManagersManagerIdNetworkProtocolGet(ctx, manager)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response: %+v", l)
}
