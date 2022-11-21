package market

import "github.com/iaping/go-okx/rest/api"

func NewGetPlatform24Volume() (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/platform-24-volume",
		Method: api.MethodGet,
	}, &GetPlatform24VolumeResponse{}
}

type GetPlatform24VolumeResponse struct {
	api.Response
	Data []Platform24Volume `json:"data"`
}

type Platform24Volume struct {
	VolUsd string `json:"volUsd"`
	VolCny string `json:"volCny"`
	Ts     int64  `json:"ts,string"`
}
