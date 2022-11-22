package main

import (
	"log"

	"github.com/iaping/go-okx/ws"
	"github.com/iaping/go-okx/ws/public"
)

func main() {
	args := &ws.Args{
		Channel: "candle1m",
		InstId:  "BTC-USDT",
	}
	handler := func(c public.EventKline) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeKline(args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
