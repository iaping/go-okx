package account

import "github.com/iaping/go-okx/rest/api"

func NewGetMaxAvailSize(param *GetMaxAvailSizeParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/max-avail-size",
		Method: api.MethodGet,
		Param:  param,
	}, &GetMaxAvailSizeResponse{}
}

type GetMaxAvailSizeParam struct {
	InstId     string `url:"instId"`
	TdMode     string `url:"tdMode"`
	Ccy        string `url:"ccy,omitempty"`
	Px         string `url:"px,omitempty"`
	ReduceOnly string `url:"reduceOnly,omitempty"`
}

type GetMaxAvailSizeResponse struct {
	api.Response
	Data []MaxSize `json:"data"`
}

type MaxAvailSize struct {
	InstId    string `json:"instId"`
	AvailBuy  string `json:"availBuy"`
	AvailSell string `json:"availSell"`
}
