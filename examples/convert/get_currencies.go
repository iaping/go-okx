package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/convert"
)

func main() {
	req, resp := convert.NewGetCurrencies()
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*convert.GetCurrenciesResponse))
}
