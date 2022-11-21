package account

import "github.com/iaping/go-okx/rest/api"

func NewGetBalance(param *GetBalanceParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/balance",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBalanceResponse{}
}

type GetBalanceParam struct {
	Ccy string `url:"ccy,omitempty"`
}

type GetBalanceResponse struct {
	api.Response
	Data []Balance `json:"data"`
}

type Balance struct {
	AdjEq       string          `json:"adjEq"`
	Imr         string          `json:"imr"`
	IsoEq       string          `json:"isoEq"`
	MgnRatio    string          `json:"mgnRatio"`
	Mmr         string          `json:"mmr"`
	NotionalUsd string          `json:"notionalUsd"`
	OrdFroz     string          `json:"ordFroz"`
	TotalEq     string          `json:"totalEq"`
	UTime       int64           `json:"uTime,string"`
	Details     []BalanceDetail `json:"details"`
}

type BalanceDetail struct {
	AvailBal      string `json:"availBal"`
	AvailEq       string `json:"availEq"`
	CashBal       string `json:"cashBal"`
	Ccy           string `json:"ccy"`
	CrossLiab     string `json:"crossLiab"`
	DisEq         string `json:"disEq"`
	Eq            string `json:"eq"`
	EqUsd         string `json:"eqUsd"`
	FrozenBal     string `json:"FrozenBal"`
	Interest      string `json:"interest"`
	IsoEq         string `json:"isoEq"`
	IsoLiab       string `json:"isoLiab"`
	IsoUpl        string `json:"isoUpl"`
	Liab          string `json:"liab"`
	MaxLoan       string `json:"maxLoan"`
	MgnRatio      string `json:"mgnRatio"`
	NotionalLever string `json:"notionalLever"`
	OrdFrozen     string `json:"ordFrozen"`
	StgyEq        string `json:"stgyEq"`
	Twap          string `json:"twap"`
	UTime         int64  `json:"uTime,string"`
	Upl           string `json:"upl"`
	UplLiab       string `json:"uplLiab"`
}
