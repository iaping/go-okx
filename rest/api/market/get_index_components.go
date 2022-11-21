package market

import "github.com/iaping/go-okx/rest/api"

func NewGetIndexComponents(param *GetIndexComponentsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/index-components",
		Method: api.MethodGet,
		Param:  param,
	}, &GetIndexComponentsResponse{}
}

type GetIndexComponentsParam struct {
	Index string `url:"index"`
}

type GetIndexComponentsResponse struct {
	api.Response
	Data IndexComponents `json:"data"`
}

type IndexComponents struct {
	Index      string      `json:"index"`
	Last       string      `json:"last"`
	Ts         int64       `json:"ts,string"`
	Components []Component `json:"components"`
}

type Component struct {
	Exch   string `json:"exch"`
	Symbol string `json:"symbol"`
	SymPx  string `json:"symPx"`
	Wgt    string `json:"wgt"`
	CnvPx  string `json:"cnvPx"`
}
