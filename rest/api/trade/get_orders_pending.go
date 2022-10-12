package trade

import (
	"go-okx/rest/api"
)

func NewGetOrdersPending(param *GetOrdersPendingParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-pending",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrdersPendingResponse{}
}

type GetOrdersPendingParam struct {
	InstType string `url:"instType,omitempty"`
	Uly      string `url:"uly,omitempty"`
	InstId   string `url:"instId,omitempty"`
	OrdType  string `url:"ordType,omitempty"`
	State    string `url:"state,omitempty"`
	After    string `url:"after,omitempty"`
	Before   string `url:"before,omitempty"`
	Limit    string `url:"limit,omitempty"`
}

type GetOrdersPendingResponse struct {
	api.Response
	Data []struct {
		InstType        string `json:"instType"`
		InstId          string `json:"instId"`
		TgtCcy          string `json:"tgtCcy"`
		Ccy             string `json:"ccy"`
		OrdId           string `json:"ordId"`
		ClOrdId         string `json:"clOrdId"`
		Tag             string `json:"tag"`
		Px              string `json:"px"`
		Sz              string `json:"sz"`
		Pnl             string `json:"pnl"`
		OrdType         string `json:"ordType"`
		Side            string `json:"side"`
		PosSide         string `json:"posSide"`
		TdMode          string `json:"tdMode"`
		AccFillSz       string `json:"accFillSz"`
		FillPx          string `json:"fillPx"`
		TradeId         string `json:"tradeId"`
		FillSz          string `json:"fillSz"`
		FillTime        string `json:"fillTime"`
		AvgPx           string `json:"avgPx"`
		State           string `json:"state"`
		Lever           string `json:"lever"`
		TpTriggerPx     string `json:"tpTriggerPx"`
		TpTriggerPxType string `json:"tpTriggerPxType"`
		SlTriggerPx     string `json:"slTriggerPx"`
		SlTriggerPxType string `json:"slTriggerPxType"`
		SlOrdPx         string `json:"slOrdPx"`
		FeeCcy          string `json:"feeCcy"`
		Fee             string `json:"fee"`
		RebateCcy       string `json:"rebateCcy"`
		Source          string `json:"source"`
		Rebate          string `json:"rebate"`
		Category        string `json:"category"`
		UTime           int64  `json:"uTime,string"`
		CTime           int64  `json:"cTime,string"`
	} `json:"data"`
}
