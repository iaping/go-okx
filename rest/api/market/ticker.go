package market

type Ticker struct {
	InstType  string  `json:"instType"`
	InstId    string  `json:"instId"`
	Last      string  `json:"last"`
	LastSz    string  `json:"lastSz"`
	AskPx     string  `json:"askPx"`
	AskSz     string  `json:"askSz"`
	BidPx     string  `json:"bidPx"`
	BidSz     string  `json:"bidSz"`
	Open24h   float64 `json:"open24h,string"`
	High24h   float64 `json:"high24h,string"`
	Low24h    float64 `json:"low24h,string"`
	VolCcy24h float64 `json:"volCcy24h,string"`
	Vol24h    float64 `json:"vol24h,string"`
	SodUtc0   float64 `json:"sodUtc0,string"`
	SodUtc8   float64 `json:"sodUtc8,string"`
	Ts        int64   `json:"ts,string"`
}
