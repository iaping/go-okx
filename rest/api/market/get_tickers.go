package market

import "github.com/iaping/go-okx/rest/api"

func NewGetTickers(param *GetTickersParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/tickers",
		Method: api.MethodGet,
		Param:  param,
	}, &GetTickersResponse{}
}

type GetTickersParam struct {
	InstType string `url:"instType"`
	Uly      string `url:"uly,omitempty"`
}

type GetTickersResponse struct {
	api.Response
	Data []struct {
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
	} `json:"data"`
}
