package main

import (
	"log"

	"github.com/iaping/go-okx/examples/rest"
	"github.com/iaping/go-okx/rest/api/asset"
)

func main() {
	param := &asset.GetAssetValuationParam{}
	req, resp := asset.NewGetAssetValuation(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*asset.GetAssetValuationResponse))
}
