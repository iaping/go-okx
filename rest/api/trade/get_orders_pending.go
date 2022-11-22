package trade

import (
	"github.com/iaping/go-okx/rest/api"
)

func NewGetOrdersPending(param *GetOrdersQueryParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-pending",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrderResponse{}
}

type GetOrdersQueryParam struct {
	api.PagingParam
	InstType string `url:"instType,omitempty"`
	Uly      string `url:"uly,omitempty"`
	InstId   string `url:"instId,omitempty"`
	OrdType  string `url:"ordType,omitempty"`
	State    string `url:"state,omitempty"`
	Category string `url:"category,omitempty"`
}
