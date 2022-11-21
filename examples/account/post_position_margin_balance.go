package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api"
	"github.com/iaping/go-okx/rest/api/account"
)

func main() {
	param := &account.PostPositionMarginBalanceParam{
		InstId:  "BTC-USDT-SWAP",
		PosSide: api.PosSideNet,
		Type:    api.TypeReduce,
		Amt:     "1",
	}
	req, resp := account.NewPostPositionMarginBalance(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.PostPositionMarginBalanceResponse))
}
