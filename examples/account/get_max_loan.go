package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api"
	"github.com/iaping/go-okx/rest/api/account"
)

func main() {
	param := &account.GetMaxLoanParam{
		InstId:  "BTC-USDT",
		MgnMode: api.MgnModeCross,
		MgnCcy:  "USDT",
	}
	req, resp := account.NewGetMaxLoan(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetMaxLoanResponse))
}
