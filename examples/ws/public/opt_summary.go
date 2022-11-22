package main

import (
	"log"

	"github.com/iaping/go-okx/ws/public"
)

func main() {
	handler := func(c public.EventOptSummary) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeOptSummary("BTC-USDT", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
