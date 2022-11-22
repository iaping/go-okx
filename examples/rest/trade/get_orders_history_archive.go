package main

import (
	"log"

	"github.com/iaping/go-okx/examples/rest"
	"github.com/iaping/go-okx/rest/api"
	"github.com/iaping/go-okx/rest/api/trade"
)

func main() {
	param := &trade.GetOrdersQueryParam{
		InstType: api.InstTypeSPOT,
	}
	req, resp := trade.NewGetOrdersHistoryArchive(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetOrderResponse))
}
