package ws

const (
	OpSubscribe   = "subscribe"
	OpUnsubscribe = "unsubscribe"
	OpLogin       = "login"
)

type Request struct {
	Op   string      `json:"op"`
	Args interface{} `json:"args"`
}

// new request for subscribe
func NewRequestSubscribe(args interface{}) *Request {
	return NewRequest(OpSubscribe, args)
}

// new request for login
func NewRequestLogin(args interface{}) *Request {
	return NewRequest(OpLogin, args)
}

// new request
func NewRequest(op string, args interface{}) *Request {
	return &Request{
		Op:   op,
		Args: []interface{}{args},
	}
}
