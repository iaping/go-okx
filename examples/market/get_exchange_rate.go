package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/market"
)

func main() {
	req, resp := market.NewGetExchangeRate()
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetExchangeRateResponse))
}
