package market

import "github.com/iaping/go-okx/rest/api"

func NewGetExchangeRate() (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/exchange-rate",
		Method: api.MethodGet,
	}, &GetExchangeRateResponse{}
}

type GetExchangeRateResponse struct {
	api.Response
	Data []ExchangeRate `json:"data"`
}

type ExchangeRate struct {
	UsdCny string `json:"usdCny"`
}
