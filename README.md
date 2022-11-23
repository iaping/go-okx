# go-okx
golang sdk for https://www.okx.com/docs-v5

## Installation

go get：
```shell
go get github.com/iaping/go-okx
```

## Example
all examples are in the folder examples

## Rest api example code
```go
package main

import (
	"log"

	"github.com/iaping/go-okx/rest"
	"github.com/iaping/go-okx/rest/api/account"
)

func main() {
	auth := common.NewAuth("your apikey", "your key", "your passphrase", false)
	client := rest.New("", auth, nil)
	param := &account.GetBalanceParam{}
	req, resp := account.NewGetBalance(param)
	if err := client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetBalanceResponse))
}
```

## Public websocket example code
```go
package main

import (
	"log"

	"github.com/iaping/go-okx/ws/public"
)

func main() {
	handler := func(c public.EventTickers) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeTickers("BTC-USDT", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
```

## Private websocket example code
```go
package main

import (
	"log"

	"github.com/iaping/go-okx/common"
	"github.com/iaping/go-okx/ws"
	"github.com/iaping/go-okx/ws/private"
)

func main() {
	auth := common.NewAuth("your apikey", "your key", "your passphrase", false)
	args := &ws.Args{
		InstType: "SPOT",
	}
	handler := func(c private.EventOrders) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := private.SubscribeOrders(args, auth, handler, handlerError); err != nil {
		panic(err)
	}
	select {}
}
```

## 提示
缺失的接口会慢慢完善（完善的进度看Star）。有部分接口只是调通了，没有真实数据测试（穷）。有问题欢迎Issues！！！