package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/trade"
)

func main() {
	param := []*trade.PostAmendOrderParam{
		&trade.PostAmendOrderParam{
			InstId: "OKB-USDT",
			OrdId:  "123456789",
		},
		&trade.PostAmendOrderParam{
			InstId: "XRP-USDT",
			OrdId:  "123456789",
		},
	}
	req, resp := trade.NewPostAmendBatchOrders(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostAmendOrderResponse))
}
