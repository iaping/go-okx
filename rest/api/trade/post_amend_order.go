package trade

import "github.com/iaping/go-okx/rest/api"

func NewPostAmendOrder(param *PostAmendOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/amend-order",
		Method: api.MethodPost,
		Param:  param,
	}, &PostAmendOrderResponse{}
}

type PostAmendOrderParam struct {
	InstId    string `json:"instId"`
	CxlOnFail bool   `json:"cxlOnFail,omitempty"`
	OrdId     string `json:"ordId,omitempty"`
	ClOrdId   string `json:"clOrdId,omitempty"`
	ReqId     string `json:"reqId,omitempty"`
	NewSz     string `json:"newSz,omitempty"`
	NewPx     string `json:"newPx,omitempty"`
}

type PostAmendOrderResponse struct {
	api.Response
	Data []AmendOrder `json:"data"`
}

type AmendOrder struct {
	OrdId   string `json:"ordId"`
	ClOrdId string `json:"clOrdId"`
	ReqId   string `json:"reqId"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}
