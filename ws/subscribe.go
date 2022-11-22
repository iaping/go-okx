package ws

type Handler func(message []byte)
type HandlerError func(err error)

type Subscribe struct {
	Request      Request
	Response     Response
	Handler      func(message []byte)
	HandlerError HandlerError
}

// new subscribe
func NewSubscribe(args interface{}, handler Handler, handlerError HandlerError) *Subscribe {
	return &Subscribe{
		Request:      NewRequestSubscribe(args),
		Handler:      handler,
		HandlerError: handlerError,
	}
}
