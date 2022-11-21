package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/trade"
)

func main() {
	param := &trade.PostAmendOrderParam{
		InstId: "OKB-USDT",
		OrdId:  "123456789",
	}
	req, resp := trade.NewPostAmendOrder(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostAmendOrderResponse))
}
