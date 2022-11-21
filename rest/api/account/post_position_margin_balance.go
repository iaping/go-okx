package account

import "github.com/iaping/go-okx/rest/api"

func NewPostPositionMarginBalance(param *PostPositionMarginBalanceParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/position/margin-balance",
		Method: api.MethodPost,
		Param:  param,
	}, &PostPositionMarginBalanceResponse{}
}

type PostPositionMarginBalanceParam struct {
	InstId    string `json:"instId"`
	PosSide   string `json:"posSide"`
	Type      string `json:"type"`
	Amt       string `json:"amt"`
	Ccy       string `json:"ccy,omitempty"`
	Auto      bool   `json:"auto,omitempty"`
	LoanTrans bool   `json:"loanTrans,omitempty"`
}

type PostPositionMarginBalanceResponse struct {
	api.Response
	Data []PositionMarginBalance `json:"data"`
}

type PositionMarginBalance struct {
	InstId   string `json:"instId"`
	PosSide  string `json:"posSide"`
	Amt      string `json:"amt"`
	Type     string `json:"type"`
	Leverage string `json:"leverage"`
	Ccy      string `json:"ccy"`
}
