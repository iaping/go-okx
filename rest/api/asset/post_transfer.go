package asset

import "github.com/iaping/go-okx/rest/api"

func NewPostTransfer(param *PostTransferParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/transfer",
		Method: api.MethodPost,
		Param:  param,
	}, &PostTransferResponse{}
}

type PostTransferParam struct {
	Ccy       string `json:"ccy"`
	Amt       string `json:"amt"`
	From      string `json:"from"`
	To        string `json:"to"`
	SubAcct   string `json:"subAcct,omitempty"`
	Type      string `json:"type,omitempty"`
	LoanTrans bool   `json:"loanTrans,omitempty"`
}

type PostTransferResponse struct {
	api.Response
	Data []Transfer `json:"data"`
}

type Transfer struct {
	TransId string `json:"transId"`
	Ccy     string `json:"ccy"`
	From    string `json:"from"`
	Amt     string `json:"amt"`
	To      string `json:"to"`
}
