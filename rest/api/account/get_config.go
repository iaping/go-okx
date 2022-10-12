package account

import "go-okx/rest/api"

func NewGetConfig() (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/config",
		Method: api.MethodGet,
	}, &GetBalanceResponse{}
}

type GetConfigResponse struct {
	api.Response
	Data []struct {
		Uid        uint64 `json:"uid,string"`
		AcctLv     int    `json:"acctLv,string"`
		PosMode    string `json:"posMode"`
		AutoLoan   bool   `json:"autoLoan"`
		GreeksType string `json:"greeksType"`
		Level      string `json:"level"`
		LevelTmp   string `json:"levelTmp"`
		CtIsoMode  string `json:"ctIsoMode"`
		MgnIsoMode string `json:"mgnIsoMode"`
	} `json:"data"`
}
