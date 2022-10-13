# go-okx
golang sdk for https://www.okx.com/docs-v5

## Installation

go get：
```shell
go get github.com/iaping/go-okx
```

## Example
all examples are in the folder examples

## Example code
```go
package main

import (
	"log"

	"github.com/iaping/go-okx/rest"
	"github.com/iaping/go-okx/rest/api/account"
)

func main() {
	client := rest.New("", "your apikey", "your key", "your passphrase", false, nil)
	param := &account.GetBalanceParam{}
	req, resp := account.NewGetBalance(param)
	if err := client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetBalanceResponse))
}
```