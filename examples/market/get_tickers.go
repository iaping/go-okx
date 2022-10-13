package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api"
	"github.com/iaping/go-okx/rest/api/market"
)

func main() {
	param := &market.GetTickersParam{
		InstType: api.InstTypeSPOT,
	}
	req, resp := market.NewGetTickers(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetTickersResponse))
}
