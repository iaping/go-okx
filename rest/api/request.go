package api

const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

type IRequest interface {
	GetPath() string
	GetMethod() string
	GetParam() interface{}
	IsPost() bool
}

type Request struct {
	Path   string
	Method string
	Param  interface{}
}

func (r Request) GetPath() string {
	return r.Path
}

func (r Request) GetMethod() string {
	return r.Method
}

func (r Request) GetParam() interface{} {
	return r.Param
}

func (r Request) IsPost() bool {
	return r.Method == MethodPost
}
