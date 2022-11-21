package account

import "github.com/iaping/go-okx/rest/api"

func NewPostSetPositionMode(param *PostSetPositionModeParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/set-position-mode",
		Method: api.MethodPost,
		Param:  param,
	}, &PostSetPositionModeResponse{}
}

type PostSetPositionModeParam struct {
	PosMode string `json:"posMode"`
}

type PostSetPositionModeResponse struct {
	api.Response
	Data []SetPositionMode `json:"data"`
}

type SetPositionMode struct {
	PosMode string `json:"posMode"`
}
