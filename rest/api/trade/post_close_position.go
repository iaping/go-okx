package trade

import "github.com/iaping/go-okx/rest/api"

func NewPostClosePosition(param *PostClosePositionParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/close-position",
		Method: api.MethodPost,
		Param:  param,
	}, &PostClosePositionResponse{}
}

type PostClosePositionParam struct {
	InstId  string `json:"instId"`
	PosSide string `json:"posSide,omitempty"`
	MgnMode string `json:"mgnMode"`
	Ccy     string `json:"ccy,omitempty"`
	AutoCxl bool   `json:"autoCxl,omitempty"`
}

type PostClosePositionResponse struct {
	api.Response
	Data []ClosePosition `json:"data"`
}

type ClosePosition struct {
	InstId  string `json:"instId"`
	PosSide string `json:"posSide"`
}
