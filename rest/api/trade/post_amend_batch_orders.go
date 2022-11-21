package trade

import "github.com/iaping/go-okx/rest/api"

func NewPostAmendBatchOrders(param []*PostAmendOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/amend-batch-orders",
		Method: api.MethodPost,
		Param:  param,
	}, &PostAmendOrderResponse{}
}
