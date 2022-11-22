package public

import (
	"encoding/json"

	"github.com/iaping/go-okx/ws"
)

type HandlerPriceLimit func(EventPriceLimit)

type EventPriceLimit struct {
	Arg  ws.Args      `json:"arg"`
	Data []PriceLimit `json:"data"`
}

type PriceLimit struct {
	InstId  string `json:"instId"`
	BuyLmt  string `json:"buyLmt"`
	SellLmt string `json:"sellLmt"`
	Ts      int64  `json:"ts,string"`
}

// default subscribe
func SubscribePriceLimit(instId string, handler HandlerPriceLimit, handlerError ws.HandlerError, simulated bool) error {
	args := &ws.Args{
		Channel: "price-limit",
		InstId:  instId,
	}

	h := func(message []byte) {
		var event EventPriceLimit
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
