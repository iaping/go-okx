package main

import (
	"log"

	"github.com/iaping/go-okx/examples/rest"
	"github.com/iaping/go-okx/rest/api"
	"github.com/iaping/go-okx/rest/api/account"
)

func main() {
	param := &account.GetMaxSizeParam{
		InstId: "BTC-USDT",
		TdMode: api.TdModeCross,
	}
	req, resp := account.NewGetMaxSize(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetMaxSizeResponse))
}
