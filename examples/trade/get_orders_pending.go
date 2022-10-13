package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/trade"
)

func main() {
	param := &trade.GetOrdersPendingParam{}
	req, resp := trade.NewGetOrdersPending(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetOrdersPendingResponse))
}
