package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/asset"
)

func main() {
	param := &asset.PostTransferParam{
		Ccy:  "USDT",
		Amt:  "1",
		From: "18",
		To:   "6",
	}
	req, resp := asset.NewPostTransfer(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*asset.PostTransferResponse))
}
