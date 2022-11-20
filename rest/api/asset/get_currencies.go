package asset

import "github.com/iaping/go-okx/rest/api"

func NewGetCurrencies(param *GetCurrenciesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/currencies",
		Method: api.MethodGet,
		Param:  param,
	}, &GetCurrenciesResponse{}
}

type GetCurrenciesParam struct {
	Ccy string `url:"ccy,omitempty"`
}

type GetCurrenciesResponse struct {
	api.Response
	Data []Currencies `json:"data"`
}

type Currencies struct {
	Ccy                  string `json:"ccy"`
	Name                 string `json:"name"`
	LogoLink             string `json:"logoLink"`
	Chain                string `json:"chain"`
	CanDep               bool   `json:"canDep"`
	CanWd                bool   `json:"canWd"`
	CanInternal          bool   `json:"canInternal"`
	MinDep               string `json:"minDep"`
	MinWd                string `json:"minWd"`
	MaxWd                string `json:"maxWd"`
	WdTickSz             string `json:"wdTickSz"`
	WdQuota              string `json:"wdQuota"`
	UsedWdQuota          string `json:"usedWdQuota"`
	MinFee               string `json:"minFee"`
	MaxFee               string `json:"maxFee"`
	MainNet              bool   `json:"mainNet"`
	NeedTag              bool   `json:"needTag"`
	MinDepArrivalConfirm string `json:"minDepArrivalConfirm"`
	MinWdUnlockConfirm   string `json:"minWdUnlockConfirm"`
	DepQuotaFixed        string `json:"depQuotaFixed"`
	UsedDepQuotaFixed    string `json:"usedDepQuotaFixed"`
}
