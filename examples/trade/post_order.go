package main

import (
	"go-okx/examples"
	"go-okx/rest/api/trade"
	"log"
)

func main() {
	param := &trade.PostOrderParam{
		InstId:  "OKB-USDT",
		TdMode:  trade.TdModeCash,
		Side:    trade.SideBuy,
		OrdType: trade.OrdTypeLimit,
		Sz:      "9",
		Px:      "5",
	}
	req, resp := trade.NewPostOrder(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostOrderResponse))
}
