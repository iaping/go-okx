package trade

import "go-okx/rest/api"

func NewPostOrder(param *PostOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/order",
		Method: api.MethodPost,
		Param:  param,
	}, &PostOrderResponse{}
}

type PostOrderParam struct {
	InstId     string `json:"instId"`
	TdMode     string `json:"tdMode"`
	Ccy        string `json:"ccy,omitempty"`
	ClOrdId    string `json:"clOrdId,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Side       string `json:"side"`
	PosSide    string `json:"posSide,omitempty"`
	OrdType    string `json:"ordType"`
	Sz         string `json:"sz"`
	Px         string `json:"px,omitempty"`
	ReduceOnly bool   `json:"reduceOnly,omitempty"`
	TgtCcy     string `json:"tgtCcy,omitempty"`
}

type PostOrderResponse struct {
	api.Response
	Data []struct {
		OrdId   string `json:"ordId"`
		ClOrdId string `json:"clOrdId"`
		Tag     string `json:"tag"`
		SCode   string `json:"sCode"`
		SMsg    string `json:"sMsg"`
	} `json:"data"`
}
