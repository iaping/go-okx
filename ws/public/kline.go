package public

import (
	"encoding/json"

	"github.com/iaping/go-okx/ws"
)

type HandlerKline func(EventKline)

type EventKline struct {
	Arg  ws.Args    `json:"arg"`
	Data [][]string `json:"data"`
}

// default subscribe
func SubscribeKline(args *ws.Args, handler HandlerKline, handlerError ws.HandlerError, simulated bool) error {
	h := func(message []byte) {
		var event EventKline
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
