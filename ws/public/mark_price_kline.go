package public

import (
	"encoding/json"

	"github.com/iaping/go-okx/ws"
)

type HandlerMarkPriceKline func(EventMarkPriceKline)

type EventMarkPriceKline struct {
	Arg  ws.Args    `json:"arg"`
	Data [][]string `json:"data"`
}

// default subscribe
func SubscribeMarkPriceKline(args *ws.Args, handler HandlerMarkPriceKline, handlerError ws.HandlerError, simulated bool) error {
	h := func(message []byte) {
		var event EventMarkPriceKline
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
