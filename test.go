package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/bmc-toolbox/redgopher/openapi"
	"github.com/davecgh/go-spew/spew"
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

	cfg := &openapi.Configuration{
		BasePath:   fmt.Sprintf("http://%s", ip),
		HTTPClient: &http.Client{Transport: tr},
	}
	client := openapi.NewAPIClient(cfg)

	ctx := context.Background()
	basicAuth := openapi.BasicAuth{UserName: user, Password: pass}
	ctx2 := context.WithValue(ctx, openapi.ContextBasicAuth, basicAuth)

	//	sessionParam := openapi.Session{UserName: user, Password: pass }
	//	sessions, _, err := client.DefaultApi.RedfishV1SessionServiceSessionsPost(ctx2, sessionParam)
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	//	spew.Dump(sessions)

	l, _, err := client.DefaultApi.RedfishV1ManagersManagerIdLogServicesGet(ctx2, "1")
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(l)
}
