package trade

import "github.com/iaping/go-okx/rest/api"

func NewPostCancelOrder(param *PostCancelOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/cancel-order",
		Method: api.MethodPost,
		Param:  param,
	}, &PostCancelOrderResponse{}
}

type PostCancelOrderParam struct {
	InstId  string `json:"instId"`
	OrdId   string `json:"ordId,omitempty"`
	ClOrdId string `json:"clOrdId,omitempty"`
}

type PostCancelOrderResponse struct {
	api.Response
	Data []CancelOrder `json:"data"`
}

type CancelOrder struct {
	OrdId   string `json:"ordId"`
	ClOrdId string `json:"clOrdId"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}
