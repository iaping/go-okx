package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/trade"
)

func main() {
	param := []*trade.PostCancelOrderParam{
		&trade.PostCancelOrderParam{
			InstId: "OKB-USDT",
			OrdId:  "515102281109434368",
		},
		&trade.PostCancelOrderParam{
			InstId: "XRP-USDT",
			OrdId:  "515102281109434369",
		},
	}
	req, resp := trade.NewPostCancelBatchOrders(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostCancelOrderResponse))
}
