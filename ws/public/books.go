package public

import (
	"encoding/json"

	"github.com/iaping/go-okx/ws"
)

type HandlerBooks func(EventBooks)

type EventBooks struct {
	Arg    ws.Args `json:"arg"`
	Data   []Book  `json:"data"`
	Action string  `json:"action"`
}

type Book struct {
	Asks     [][]string `json:"asks"`
	Bids     [][]string `json:"bids"`
	Ts       int64      `json:"ts,string"`
	Checksum int64      `json:"checksum"`
}

// default subscribe
func SubscribeBooks(args *ws.Args, handler HandlerBooks, handlerError ws.HandlerError, simulated bool) error {
	h := func(message []byte) {
		var event EventBooks
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
