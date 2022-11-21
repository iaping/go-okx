package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/market"
)

func main() {
	param := &market.GetCandlesParam{
		InstId: "ETH-USDT",
	}
	req, resp := market.NewGetIndexCandles(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetCandlesResponse))
}
