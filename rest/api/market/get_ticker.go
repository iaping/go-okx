package market

import "github.com/iaping/go-okx/rest/api"

func NewGetTicker(param *GetTickerParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/ticker",
		Method: api.MethodGet,
		Param:  param,
	}, &GetTickerResponse{}
}

type GetTickerParam struct {
	InstId string `url:"instId"`
}

type GetTickerResponse struct {
	api.Response
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
