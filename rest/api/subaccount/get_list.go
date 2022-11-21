package subaccount

import "github.com/iaping/go-okx/rest/api"

func NewGetList(param *GetListParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/users/subaccount/list",
		Method: api.MethodGet,
		Param:  param,
	}, &GetListResponse{}
}

type GetListParam struct {
	Enable  bool   `url:"enable,omitempty"`
	SubAcct string `url:"subAcct,omitempty"`
	After   int64  `url:"after,omitempty"`
	Before  int64  `url:"before,omitempty"`
	Limit   int    `url:"limit,omitempty"`
}

type GetListResponse struct {
	api.Response
	Data []Account `json:"data"`
}

type Account struct {
	Type    string `json:"type"`
	Enable  bool   `json:"enable"`
	SubAcct string `json:"subAcct"`
	Label   string `json:"label"`
	Mobile  string `json:"mobile"`
	GAuth   bool   `json:"gAuth"`
	Ts      int64  `json:"ts,string"`
}
