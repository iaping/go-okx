package trade

import "github.com/iaping/go-okx/rest/api"

func NewPostBatchOrders(param []*PostOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/batch-orders",
		Method: api.MethodPost,
		Param:  param,
	}, &PostOrderResponse{}
}
