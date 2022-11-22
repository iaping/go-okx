package main

import (
	"log"

	"github.com/iaping/go-okx/examples/rest"
	"github.com/iaping/go-okx/rest/api/convert"
)

func main() {
	param := &convert.GetHistoryParam{}
	req, resp := convert.NewGetHistory(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*convert.GetHistoryResponse))
}
