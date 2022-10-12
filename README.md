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
    "go-okx/rest"
    "go-okx/rest/api/account"
    "log"
)

func main() {
    client = rest.New("", "your api key", "you key", "your passphrase", false, nil)
    param := &account.GetBalanceParam{}
    req, resp := account.NewGetBalance(param)
    if err := client.Do(req, resp); err != nil {
        panic(err)
    }
    log.Print(req, resp.(*account.GetBalanceResponse))
}
```