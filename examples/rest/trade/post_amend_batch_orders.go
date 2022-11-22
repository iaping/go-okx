package main

import (
	"log"

	"github.com/iaping/go-okx/examples/rest"
	"github.com/iaping/go-okx/rest/api/trade"
)

func main() {
	param := []*trade.PostAmendOrderParam{
		&trade.PostAmendOrderParam{
			InstId: "OKB-USDT",
			OrdId:  "515102546340442112",
			NewSz:  "1.8",
		},
		&trade.PostAmendOrderParam{
			InstId: "XRP-USDT",
			OrdId:  "515102546340442113",
			NewSz:  "2.8",
		},
	}
	req, resp := trade.NewPostAmendBatchOrders(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostAmendOrderResponse))
}
