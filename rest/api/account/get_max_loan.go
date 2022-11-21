package account

import "github.com/iaping/go-okx/rest/api"

func NewGetMaxLoan(param *GetMaxLoanParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/max-loan",
		Method: api.MethodGet,
		Param:  param,
	}, &GetMaxLoanResponse{}
}

type GetMaxLoanParam struct {
	InstId  string `url:"instId"`
	MgnMode string `url:"mgnMode"`
	MgnCcy  string `url:"mgnCcy,omitempty"`
}

type GetMaxLoanResponse struct {
	api.Response
	Data []MaxLoan `json:"data"`
}

type MaxLoan struct {
	InstId  string `json:"instId"`
	MgnMode string `json:"mgnMode"`
	MgnCcy  string `json:"mgnCcy"`
	MaxLoan string `json:"maxLoan"`
	Ccy     string `json:"ccy"`
	Side    string `json:"side"`
}
