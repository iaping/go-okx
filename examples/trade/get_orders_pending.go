package main

import (
	"go-okx/examples"
	"go-okx/rest/api/trade"
	"log"
)

func main() {
	param := &trade.GetOrdersPendingParam{}
	req, resp := trade.NewGetOrdersPending(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetOrdersPendingResponse))
}
