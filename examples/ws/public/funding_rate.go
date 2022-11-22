package main

import (
	"log"

	"github.com/iaping/go-okx/ws/public"
)

func main() {
	handler := func(c public.EventFundingRate) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeFundingRate("BTC-USDT-SWAP", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
