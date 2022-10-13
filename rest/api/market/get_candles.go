package market

import "github.com/iaping/go-okx/rest/api"

func NewGetCandles(param *GetCandlesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/candles",
		Method: api.MethodGet,
		Param:  param,
	}, &GetCandlesResponse{}
}

type GetCandlesParam struct {
	InstId string `url:"instId"`
	Bar    string `url:"bar,omitempty"`
	After  int64  `url:"after,omitempty"`
	Before int64  `url:"before,omitempty"`
	Limit  int    `url:"limit,omitempty"`
}

type GetCandlesResponse struct {
	api.Response
	Data [][]string `json:"data"`
}
