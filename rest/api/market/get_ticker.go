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
