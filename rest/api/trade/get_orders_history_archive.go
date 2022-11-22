package trade

import (
	"github.com/iaping/go-okx/rest/api"
)

func NewGetOrdersHistoryArchive(param *GetOrdersQueryParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-history-archive",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrderResponse{}
}
