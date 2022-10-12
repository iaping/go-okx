package main

import (
	"go-okx/examples"
	"go-okx/rest/api/account"
	"log"
)

func main() {
	param := &account.GetBalanceParam{}
	req, resp := account.NewBalance(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp)
}
