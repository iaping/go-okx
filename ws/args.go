package ws

type Args struct {
	Channel    string `json:"channel"`
	InstId     string `json:"instId,omitempty"`
	InstType   string `json:"instType,omitempty"`
	InstFamily string `json:"instFamily,omitempty"`
}
