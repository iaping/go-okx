package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/account"
)

func main() {
	param := &account.GetBalanceParam{}
	req, resp := account.NewGetBalance(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetBalanceResponse))
}
