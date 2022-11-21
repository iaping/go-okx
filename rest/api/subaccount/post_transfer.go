package subaccount

import "github.com/iaping/go-okx/rest/api"

func NewPostTransfer(param *PostTransferParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/subaccount/transfer",
		Method: api.MethodPost,
		Param:  param,
	}, &PostTransferResponse{}
}

type PostTransferParam struct {
	Ccy            string `json:"ccy"`
	Amt            string `json:"amt"`
	From           string `json:"from"`
	To             string `json:"to"`
	FromSubAccount string `json:"fromSubAccount"`
	ToSubAccount   string `json:"toSubAccount"`
	LoanTrans      bool   `json:"loanTrans,omitempty"`
}

type PostTransferResponse struct {
	api.Response
	Data []Transfer `json:"data"`
}

type Transfer struct {
	TransId string `json:"transId"`
}
