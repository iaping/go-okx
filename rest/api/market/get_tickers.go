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
	Data []Ticker `json:"data"`
}
