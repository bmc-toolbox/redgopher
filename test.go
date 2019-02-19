package main

import (
	"context"
	"log"

	"github.com/bmc-toolbox/redgopher/openapi"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	cfg := &openapi.Configuration{
		BasePath: "http://127.0.0.1",
	}
	client := openapi.NewAPIClient(cfg)

	ctx := context.Background()

	accounts, _, err := client.DefaultApi.RedfishV1AccountServiceAccountsGet(ctx)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(accounts)
}
