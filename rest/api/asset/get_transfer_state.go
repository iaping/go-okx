package asset

import "github.com/iaping/go-okx/rest/api"

func NewGetTransferState(param *GetTransferStateParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/transfer-state",
		Method: api.MethodGet,
		Param:  param,
	}, &GetTransferStateResponse{}
}

type GetTransferStateParam struct {
	TransId string `url:"transId"`
	Type    string `url:"type,omitempty"`
}

type GetTransferStateResponse struct {
	api.Response
	Data []TransferState `json:"data"`
}

type TransferState struct {
	Type     string `json:"type"`
	TransId  string `json:"transId"`
	Ccy      string `json:"ccy"`
	From     string `json:"from"`
	Amt      string `json:"amt"`
	To       string `json:"to"`
	SubAcct  string `json:"subAcct"`
	InstId   string `json:"instId"`
	ToInstId string `json:"toInstId"`
	State    string `json:"state"`
}
