package trade

import (
	"github.com/iaping/go-okx/rest/api"
)

func NewGetOrder(param *GetOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/order",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrderResponse{}
}

type GetOrderParam struct {
	InstId  string `url:"instId"`
	OrdId   string `url:"ordId,omitempty"`
	ClOrdId string `url:"clOrdId,omitempty"`
}

type GetOrderResponse struct {
	api.Response
	Data []OrderDetail `json:"data"`
}
