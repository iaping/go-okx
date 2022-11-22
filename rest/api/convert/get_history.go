package convert

import "github.com/iaping/go-okx/rest/api"

func NewGetHistory(param *GetHistoryParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/convert/history",
		Method: api.MethodGet,
		Param:  param,
	}, &GetHistoryResponse{}
}

type GetHistoryParam struct {
	api.PagingParam
}

type GetHistoryResponse struct {
	api.Response
	Data []History `json:"data"`
}

type History struct {
	TradeId     string `json:"tradeId"`
	State       string `json:"state"`
	InstId      string `json:"instId"`
	BaseCcy     string `json:"baseCcy"`
	QuoteCcy    string `json:"quoteCcy"`
	Side        string `json:"side"`
	FillPx      string `json:"fillPx"`
	FillBaseSz  string `json:"fillBaseSz"`
	FillQuoteSz string `json:"fillQuoteSz"`
	Ts          int64  `json:"ts,string"`
}
