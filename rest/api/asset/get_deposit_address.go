package asset

import "github.com/iaping/go-okx/rest/api"

func NewGetDepositAddress(param *GetDepositAddressParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/deposit-address",
		Method: api.MethodGet,
		Param:  param,
	}, &GetDepositAddressResponse{}
}

type GetDepositAddressParam struct {
	Ccy string `url:"ccy"`
}

type GetDepositAddressResponse struct {
	api.Response
	Data []DepositAddress `json:"data"`
}

type DepositAddress struct {
	Addr     string      `json:"addr"`
	Tag      string      `json:"tag"`
	Memo     string      `json:"memo"`
	PmtId    string      `json:"pmtId"`
	AddrEx   interface{} `json:"addrEx"`
	Ccy      string      `json:"ccy"`
	Chain    string      `json:"chain"`
	To       string      `json:"to"`
	Selected bool        `json:"selected"`
	CtAddr   string      `json:"ctAddr"`
}
