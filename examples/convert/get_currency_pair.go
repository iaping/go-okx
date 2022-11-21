package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/convert"
)

func main() {
	param := &convert.GetCurrencyPairParam{
		FromCcy: "BTC",
		ToCcy:   "USDT",
	}
	req, resp := convert.NewGetCurrencyPair(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*convert.GetCurrencyPairResponse))
}
