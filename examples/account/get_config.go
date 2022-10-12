package main

import (
	"go-okx/examples"
	"go-okx/rest/api/account"
	"log"
)

func main() {
	req, resp := account.NewGetConfig()
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetConfigResponse))
}
