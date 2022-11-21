package account

import "github.com/iaping/go-okx/rest/api"

func NewGetAccountPositionRisk(param *GetAccountPositionRiskParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/account-position-risk",
		Method: api.MethodGet,
		Param:  param,
	}, &GetAccountPositionRiskResponse{}
}

type GetAccountPositionRiskParam struct {
	InstType string `url:"instType,omitempty"`
}

type GetAccountPositionRiskResponse struct {
	api.Response
	Data []AccountPositionRisk `json:"data"`
}

type AccountPositionRisk struct {
	Ts      int64     `json:"ts,string"`
	AdjEq   string    `json:"adjEq"`
	BalData []BalData `json:"balData"`
	PosData []PosData `json:"posData"`
}

type BalData struct {
	Ccy   string `json:"ccy"`
	Eq    string `json:"eq"`
	DisEq string `json:"disEq"`
}

type PosData struct {
	InstType    string `json:"instType"`
	MgnMode     string `json:"mgnMode"`
	PosId       string `json:"posId"`
	InstId      string `json:"instId"`
	Pos         string `json:"pos"`
	BaseBal     string `json:"baseBal"`
	QuoteBal    string `json:"quoteBal"`
	PosSide     string `json:"posSide"`
	PosCcy      string `json:"posCcy"`
	Ccy         string `json:"ccy"`
	NotionalCcy string `json:"notionalCcy"`
	NotionalUsd string `json:"notionalUsd"`
}
