package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api"
	"github.com/iaping/go-okx/rest/api/trade"
)

func main() {
	param := &trade.GetOrdersQueryParam{
		InstType: api.InstTypeSPOT,
	}
	req, resp := trade.NewGetOrdersHistory(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetOrderResponse))
}
