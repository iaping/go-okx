package subaccount

import "github.com/iaping/go-okx/rest/api"

func NewGetBills(param *GetBillsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/subaccount/bills",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBillsResponse{}
}

type GetBillsParam struct {
	api.PagingParam
	Ccy     string `url:"ccy,omitempty"`
	Type    string `url:"type,omitempty"`
	SubAcct string `url:"subAcct,omitempty"`
}

type GetBillsResponse struct {
	api.Response
	Data []Bill `json:"data"`
}

type Bill struct {
	BillId  string `json:"billId"`
	Ccy     string `json:"ccy"`
	Amt     string `json:"amt"`
	Type    string `json:"type"`
	SubAcct string `json:"subAcct"`
	Ts      int64  `json:"ts,string"`
}
