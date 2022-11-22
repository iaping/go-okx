package api

type PagingParam struct {
	After  int64 `url:"after,omitempty"`
	Before int64 `url:"before,omitempty"`
	Limit  int   `url:"limit,omitempty"`
}
