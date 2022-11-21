package account

import "github.com/iaping/go-okx/rest/api"

func NewGetTradeFee(param *GetTradeFeeParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/trade-fee",
		Method: api.MethodGet,
		Param:  param,
	}, &GetTradeFeeResponse{}
}

type GetTradeFeeParam struct {
	InstType string `url:"instType"`
	InstId   string `url:"instId,omitempty"`
	Uly      string `url:"uly,omitempty"`
	Category string `url:"category,omitempty"`
}

type GetTradeFeeResponse struct {
	api.Response
	Data []TradeFee `json:"data"`
}

type TradeFee struct {
	Category string `json:"category"`
	Taker    string `json:"taker"`
	Maker    string `json:"maker"`
	Delivery string `json:"delivery"`
	Exercise string `json:"exercise"`
	Level    string `json:"level"`
	InstType string `json:"instType"`
	Ts       int64  `json:"ts,string"`
}
