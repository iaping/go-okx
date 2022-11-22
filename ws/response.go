package ws

import "fmt"

const (
	EventSubscribe   = OpSubscribe
	EventUnsubscribe = OpUnsubscribe
	EventError       = "error"
)

type Response struct {
	Event string      `json:"event"`
	Arg   interface{} `json:"arg"`
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
}

func (resp Response) Error() error {
	if !resp.IsError() {
		return nil
	}
	return fmt.Errorf("code: %s, msg: %s", resp.Code, resp.Msg)
}

func (resp Response) IsError() bool {
	return resp.Event == EventError
}
