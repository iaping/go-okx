package asset

import "github.com/iaping/go-okx/rest/api"

func NewGetAssetValuation(param *GetAssetValuationParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/asset-valuation",
		Method: api.MethodGet,
		Param:  param,
	}, &GetAssetValuationResponse{}
}

type GetAssetValuationParam struct {
	Ccy string `url:"ccy,omitempty"`
}

type GetAssetValuationResponse struct {
	api.Response
	Data []AssetValuation `json:"data"`
}

type AssetValuation struct {
	TotalBal string                `json:"totalBal"`
	Ts       int64                 `json:"ts,string"`
	Details  AssetValuationDetails `json:"details"`
}

type AssetValuationDetails struct {
	Funding string `json:"funding"`
	Trading string `json:"trading"`
	Classic string `json:"classic"`
	Earn    string `json:"earn"`
}
