package account

import "github.com/iaping/go-okx/rest/api"

func NewGetLeverageInfo(param *GetLeverageInfoParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/leverage-info",
		Method: api.MethodGet,
		Param:  param,
	}, &GetLeverageInfoResponse{}
}

type GetLeverageInfoParam struct {
	InstId  string `url:"instId"`
	MgnMode string `url:"mgnMode"`
}

type GetLeverageInfoResponse struct {
	api.Response
	Data []LeverageInfo `json:"data"`
}

type LeverageInfo struct {
	InstId  string `json:"instId"`
	MgnMode string `json:"mgnMode"`
	PosSide string `json:"posSide"`
	Lever   string `json:"lever"`
}
