package account

import "github.com/iaping/go-okx/rest/api"

func NewGetMaxSize(param *GetMaxSizeParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/max-size",
		Method: api.MethodGet,
		Param:  param,
	}, &GetMaxSizeResponse{}
}

type GetMaxSizeParam struct {
	InstId   string `url:"instId"`
	TdMode   string `url:"tdMode"`
	Ccy      string `url:"ccy,omitempty"`
	Px       string `url:"px,omitempty"`
	Leverage string `url:"leverage,omitempty"`
}

type GetMaxSizeResponse struct {
	api.Response
	Data []MaxSize `json:"data"`
}

type MaxSize struct {
	InstId  string `json:"instId"`
	Ccy     string `json:"ccy"`
	MaxBuy  string `json:"maxBuy"`
	MaxSell string `json:"maxSell"`
}
