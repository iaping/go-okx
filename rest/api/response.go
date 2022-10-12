package api

type IResponse interface {
	GetCode() string
	GetMessage() string
	IsOk() bool
}

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (r Response) GetCode() string {
	return r.Code
}

func (r Response) GetMessage() string {
	return r.Message
}

func (r Response) IsOk() bool {
	return r.Code == "0"
}
