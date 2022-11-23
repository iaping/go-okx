package private

import (
	"encoding/json"

	"github.com/iaping/go-okx/common"
	"github.com/iaping/go-okx/ws"
)

type HandlerBalanceAndPosition func(EventBalanceAndPosition)

type EventBalanceAndPosition struct {
	Arg  ws.Args              `json:"arg"`
	Data []BalanceAndPosition `json:"data"`
}

type BalanceAndPosition struct {
	PTime     int64     `json:"pTime,string"`
	EventType string    `json:"eventType"`
	BalData   []BalData `json:"balData"`
	PosData   []PosData `json:"posData"`
}

type BalData struct {
	Ccy     string `json:"ccy"`
	CashBal string `json:"cashBal"`
	UTime   int64  `json:"uTime,string"`
}

type PosData struct {
	PosId    string `json:"posId"`
	TradeId  string `json:"tradeId"`
	InstId   string `json:"instId"`
	InstType string `json:"instType"`
	MgnMode  string `json:"mgnMode"`
	PosSide  string `json:"posSide"`
	Pos      string `json:"pos"`
	Ccy      string `json:"ccy"`
	PosCcy   string `json:"posCcy"`
	AvgPx    string `json:"avgPx"`
	UTime    int64  `json:"uTime,string"`
}

// default subscribe
func SubscribeBalanceAndPosition(auth common.Auth, handler HandlerBalanceAndPosition, handlerError ws.HandlerError) error {
	args := &ws.Args{
		Channel: "balance_and_position",
	}

	h := func(message []byte) {
		var event EventBalanceAndPosition
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPrivate(auth).Subscribe(args, h, handlerError)
}
