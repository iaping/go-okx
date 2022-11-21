package account

import "github.com/iaping/go-okx/rest/api"

func NewGetBillsArchive(param *GetBillsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/bills-archive",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBillsResponse{}
}
