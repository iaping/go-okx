package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/convert"
)

func main() {
	param := &convert.PostTradeParam{
		BaseCcy:  "BTC",
		QuoteCcy: "USDT",
		Side:     "buy",
		Sz:       "1",
		SzCcy:    "USDT",
		QuoteId:  "123456789",
	}
	req, resp := convert.NewPostTrade(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*convert.PostTradeResponse))
}
