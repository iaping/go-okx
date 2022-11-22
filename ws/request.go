package ws

const (
	OpSubscribe   = "subscribe"
	OpUnsubscribe = "unsubscribe"
)

type Request struct {
	Op   string      `json:"op"`
	Args interface{} `json:"args"`
}

// new request for subscribe
func NewRequestSubscribe(args interface{}) Request {
	return NewRequest(OpSubscribe, args)
}

// new request
func NewRequest(op string, args interface{}) Request {
	return Request{
		Op:   op,
		Args: []interface{}{args},
	}
}
