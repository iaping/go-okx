package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/trade"
)

func main() {
	param := &trade.PostCancelOrderParam{
		InstId: "XRP-USDT",
		OrdId:  "515101542723186689",
	}
	req, resp := trade.NewPostCancelOrder(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostCancelOrderResponse))
}
