package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/market"
	"github.com/iaping/go-okx/rest/api/trade"
)

func main() {
	param := &market.GetTickersParam{
		InstType: trade.InstTypeSPOT,
	}
	req, resp := market.NewGetTickers(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetTickersResponse))
}
