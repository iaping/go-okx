package convert

import "github.com/iaping/go-okx/rest/api"

func NewPostEstimateQuote(param *PostEstimateQuoteParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/convert/estimate-quote",
		Method: api.MethodPost,
		Param:  param,
	}, &PostEstimateQuoteResponse{}
}

type PostEstimateQuoteParam struct {
	BaseCcy  string `json:"baseCcy"`
	QuoteCcy string `json:"quoteCcy"`
	Side     string `json:"side"`
	RfqSz    string `json:"rfqSz"`
	RfqSzCcy string `json:"rfqSzCcy"`
	ClQReqId string `json:"clQReqId,omitempty"`
}

type PostEstimateQuoteResponse struct {
	api.Response
	Data []EstimateQuote `json:"data"`
}

type EstimateQuote struct {
	QuoteTime int64  `json:"quoteTime,string"`
	TtlMs     int    `json:"ttlMs,string"`
	ClQReqId  string `json:"clQReqId"`
	QuoteId   string `json:"quoteId"`
	BaseCcy   string `json:"baseCcy"`
	QuoteCcy  string `json:"quoteCcy"`
	Side      string `json:"side"`
	OrigRfqSz string `json:"origRfqSz"`
	RfqSz     string `json:"rfqSz"`
	RfqSzCcy  string `json:"rfqSzCcy"`
	CnvtPx    string `json:"cnvtPx"`
	BaseSz    string `json:"baseSz"`
	QuoteSz   string `json:"quoteSz"`
}
