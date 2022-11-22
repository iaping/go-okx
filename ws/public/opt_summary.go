package public

import (
	"encoding/json"

	"github.com/iaping/go-okx/ws"
)

type HandlerOptSummary func(EventOptSummary)

type EventOptSummary struct {
	Arg  ws.Args      `json:"arg"`
	Data []OptSummary `json:"data"`
}

type OptSummary struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	Uly      string `json:"uly"`
	Delta    string `json:"delta"`
	Gamma    string `json:"gamma"`
	Theta    string `json:"theta"`
	Vega     string `json:"vega"`
	DeltaBS  string `json:"deltaBS"`
	GammaBS  string `json:"gammaBS"`
	ThetaBS  string `json:"thetaBS"`
	VegaBS   string `json:"vegaBS"`
	RealVol  string `json:"realVol"`
	BidVol   string `json:"bidVol"`
	AskVol   string `json:"askVol"`
	MarkVol  string `json:"markVol"`
	Lever    string `json:"lever"`
	FwdPx    string `json:"fwdPx"`
	Ts       int64  `json:"ts,string"`
}

// default subscribe
func SubscribeOptSummary(instFamily string, handler HandlerOptSummary, handlerError ws.HandlerError, simulated bool) error {
	args := &ws.Args{
		Channel:    "opt-summary",
		InstFamily: instFamily,
	}

	h := func(message []byte) {
		var event EventOptSummary
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
