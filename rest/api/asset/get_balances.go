package asset

import "github.com/iaping/go-okx/rest/api"

func NewGetBalances(param *GetBalancesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/balances",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBalancesResponse{}
}

type GetBalancesParam struct {
	Ccy string `url:"ccy,omitempty"`
}

type GetBalancesResponse struct {
	api.Response
	Data []Balance `json:"data"`
}

type Balance struct {
	Ccy       string `json:"ccy"`
	Bal       string `json:"bal"`
	FrozenBal string `json:"frozenBal"`
	AvailBal  string `json:"availBal"`
}
