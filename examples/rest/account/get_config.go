package main

import (
	"log"

	"github.com/iaping/go-okx/examples/rest"
	"github.com/iaping/go-okx/rest/api/account"
)

func main() {
	req, resp := account.NewGetConfig()
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetConfigResponse))
}
