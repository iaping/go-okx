package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/account"
	"github.com/iaping/go-okx/rest/api/subaccount"
)

func main() {
	param := &subaccount.GetBalancesParam{
		SubAcct: "test-01",
	}
	req, resp := subaccount.NewGetBalances(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetBalanceResponse))
}
