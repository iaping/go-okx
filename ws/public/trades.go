package public

import (
	"encoding/json"

	"github.com/iaping/go-okx/ws"
)

type HandlerTrades func(EventTrades)

type EventTrades struct {
	Arg  ws.Args `json:"arg"`
	Data []Trade `json:"data"`
}

type Trade struct {
	InstId  string `json:"instId"`
	TradeId string `json:"tradeId"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
	Side    string `json:"side"`
	Ts      int64  `json:"ts,string"`
}

// default subscribe
func SubscribeTrades(instId string, handler HandlerTrades, handlerError ws.HandlerError, simulated bool) error {
	args := &ws.Args{
		Channel: "trades",
		InstId:  instId,
	}

	h := func(message []byte) {
		var event EventTrades
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
