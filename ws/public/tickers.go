package public

import (
	"encoding/json"

	"github.com/iaping/go-okx/ws"
)

type HandlerTickers func(EventTickers)

type EventTickers struct {
	Arg  ws.Args  `json:"arg"`
	Data []Ticker `json:"data"`
}

type Ticker struct {
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24h   string `json:"open24h"`
	High24h   string `json:"high24h"`
	Low24h    string `json:"low24h"`
	VolCcy24h string `json:"volCcy24h"`
	Vol24h    string `json:"vol24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	Ts        int64  `json:"ts,string"`
}

// default subscribe
func SubscribeTickers(instId string, handler HandlerTickers, handlerError ws.HandlerError, simulated bool) error {
	args := &ws.Args{
		Channel: "tickers",
		InstId:  instId,
	}

	h := func(message []byte) {
		var event EventTickers
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
