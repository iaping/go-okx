package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/market"
)

func main() {
	param := &market.GetTradesParam{
		InstId: "ETH-USDT",
	}
	req, resp := market.NewGetTrades(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetTradesResponse))
}
