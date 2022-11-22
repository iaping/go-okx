package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api"
	"github.com/iaping/go-okx/rest/api/trade"
)

func main() {
	param := &trade.PostClosePositionParam{
		InstId:  "OKB-USDT",
		MgnMode: api.MgnModeCross,
		Ccy:     "OKB",
	}
	req, resp := trade.NewPostClosePosition(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostClosePositionResponse))
}
