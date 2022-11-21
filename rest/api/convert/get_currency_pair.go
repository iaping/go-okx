package convert

import "github.com/iaping/go-okx/rest/api"

func NewGetCurrencyPair(param *GetCurrencyPairParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/convert/currency-pair",
		Method: api.MethodGet,
		Param:  param,
	}, &GetCurrencyPairResponse{}
}

type GetCurrencyPairParam struct {
	FromCcy string `url:"fromCcy"`
	ToCcy   string `url:"toCcy"`
}

type GetCurrencyPairResponse struct {
	api.Response
	Data []CurrencyPair `json:"data"`
}

type CurrencyPair struct {
	InstId      string `json:"instId"`
	BaseCcy     string `json:"baseCcy"`
	BaseCcyMax  string `json:"baseCcyMax"`
	BaseCcyMin  string `json:"baseCcyMin"`
	QuoteCcy    string `json:"quoteCcy"`
	QuoteCcyMax string `json:"quoteCcyMax"`
	QuoteCcyMin string `json:"quoteCcyMin"`
}
