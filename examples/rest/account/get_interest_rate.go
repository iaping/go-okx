package main

import (
	"log"

	"github.com/iaping/go-okx/examples/rest"
	"github.com/iaping/go-okx/rest/api/account"
)

func main() {
	param := &account.GetInterestRateParam{
		Ccy: "USDT",
	}
	req, resp := account.NewGetInterestRate(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetInterestRateResponse))
}
