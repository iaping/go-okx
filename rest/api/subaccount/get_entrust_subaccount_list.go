package subaccount

import "github.com/iaping/go-okx/rest/api"

func NewGetEntrustSubaccountList(param *GetEntrustSubaccountListParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/users/entrust-subaccount-list",
		Method: api.MethodGet,
		Param:  param,
	}, &GetEntrustSubaccountListResponse{}
}

type GetEntrustSubaccountListParam struct {
	SubAcct string `url:"subAcct,omitempty"`
}

type GetEntrustSubaccountListResponse struct {
	api.Response
	Data []EntrustSubaccount `json:"data"`
}

type EntrustSubaccount struct {
	SubAcct string `json:"subAcct"`
}
