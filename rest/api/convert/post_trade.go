package convert

import "github.com/iaping/go-okx/rest/api"

func NewPostTrade(param *PostTradeParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/convert/trade",
		Method: api.MethodPost,
		Param:  param,
	}, &PostTradeResponse{}
}

type PostTradeParam struct {
	QuoteId  string `json:"quoteId"`
	BaseCcy  string `json:"baseCcy"`
	QuoteCcy string `json:"quoteCcy"`
	Side     string `json:"side"`
	Sz       string `json:"sz"`
	SzCcy    string `json:"szCcy"`
	ClTReqId string `json:"clTReqId,omitempty"`
}

type PostTradeResponse struct {
	api.Response
	Data []Trade `json:"data"`
}

type Trade struct {
	TradeId     string `json:"tradeId"`
	QuoteId     string `json:"quoteId"`
	ClTReqId    string `json:"clTReqId"`
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
