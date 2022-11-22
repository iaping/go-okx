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
	api.PagingParam
	InstId string `url:"instId"`
	Bar    string `url:"bar,omitempty"`
}

type GetCandlesResponse struct {
	api.Response
	Data [][]string `json:"data"`
}
