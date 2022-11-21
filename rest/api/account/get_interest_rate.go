package account

import "github.com/iaping/go-okx/rest/api"

func NewGetInterestRate(param *GetInterestRateParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/interest-rate",
		Method: api.MethodGet,
		Param:  param,
	}, &GetInterestRateResponse{}
}

type GetInterestRateParam struct {
	Ccy string `url:"ccy,omitempty"`
}

type GetInterestRateResponse struct {
	api.Response
	Data []InterestRate `json:"data"`
}

type InterestRate struct {
	InterestRate string `json:"interestRate"`
	Ccy          string `json:"ccy"`
}
