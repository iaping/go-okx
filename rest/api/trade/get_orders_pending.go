package trade

import (
	"github.com/iaping/go-okx/rest/api"
)

func NewGetOrdersPending(param *GetOrdersPendingParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-pending",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrdersPendingResponse{}
}

type GetOrdersPendingParam struct {
	InstType string `url:"instType,omitempty"`
	Uly      string `url:"uly,omitempty"`
	InstId   string `url:"instId,omitempty"`
	OrdType  string `url:"ordType,omitempty"`
	State    string `url:"state,omitempty"`
	After    int64  `url:"after,omitempty"`
	Before   int64  `url:"before,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

type GetOrdersPendingResponse struct {
	api.Response
	Data []Order `json:"data"`
}
