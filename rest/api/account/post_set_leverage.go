package account

import "github.com/iaping/go-okx/rest/api"

func NewPostSetLeverage(param *PostSetLeverageParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/set-leverage",
		Method: api.MethodPost,
		Param:  param,
	}, &PostSetLeverageResponse{}
}

type PostSetLeverageParam struct {
	InstId  string `json:"instId,omitempty"`
	Ccy     string `json:"ccy,omitempty"`
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
	PosSide string `json:"posSide,omitempty"`
}

type PostSetLeverageResponse struct {
	api.Response
	Data []SetLeverage `json:"data"`
}

type SetLeverage struct {
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
	InstId  string `json:"instId"`
	PosSide string `json:"posSide"`
}
