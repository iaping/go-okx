package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/subaccount"
)

func main() {
	param := &subaccount.PostTransferParam{
		Ccy:            "USDT",
		Amt:            "1",
		From:           "18",
		To:             "6",
		FromSubAccount: "test-01",
		ToSubAccount:   "test-02",
	}
	req, resp := subaccount.NewPostTransfer(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*subaccount.PostTransferResponse))
}
