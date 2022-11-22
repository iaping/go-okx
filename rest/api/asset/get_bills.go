package asset

import "github.com/iaping/go-okx/rest/api"

func NewGetBills(param *GetBillsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/bills",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBillsResponse{}
}

type GetBillsParam struct {
	api.PagingParam
	Ccy  string `url:"ccy,omitempty"`
	Type string `url:"type,omitempty"`
}

type GetBillsResponse struct {
	api.Response
	Data []Bill `json:"data"`
}

type Bill struct {
	BillId string `json:"billId"`
	Ccy    string `json:"ccy"`
	BalChg string `json:"balChg"`
	Bal    string `json:"bal"`
	Type   string `json:"type"`
	Ts     int64  `json:"ts,string"`
}
