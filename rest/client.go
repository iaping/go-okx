package rest

import (
	"encoding/json"
	"fmt"
	"go-okx/rest/api"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/valyala/fasthttp"
)

var (
	defaultFastHttpClient = &fasthttp.Client{
		Name:                "go-okx",
		MaxConnsPerHost:     16,
		MaxIdleConnDuration: 20 * time.Second,
		ReadTimeout:         10 * time.Second,
		WriteTimeout:        10 * time.Second,
	}
)

type Client struct {
	Host, Key, SecretKey, Passphrase string
	Simulated                        bool
	C                                *fasthttp.Client
}

// new *Client
func New(host, key, secretKey, passphrase string, simulated bool, c *fasthttp.Client) *Client {
	if host == "" {
		host = "https://www.okx.com"
	}
	if c == nil {
		c = defaultFastHttpClient
	}
	return &Client{
		Host:       host,
		Key:        key,
		SecretKey:  secretKey,
		Passphrase: passphrase,
		Simulated:  simulated,
		C:          c,
	}
}

// do request
func (c *Client) Do(req api.IRequest, resp api.IResponse) error {
	data, err := c.do(req)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}
	if !resp.IsOk() {
		return fmt.Errorf("response code: %s, message: %s", resp.GetCode(), resp.GetMessage())
	}

	return nil
}

// do request
func (c *Client) do(r api.IRequest) ([]byte, error) {
	req := c.newRequest(r)
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	if err := c.C.Do(req, resp); err != nil {
		return nil, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("http status code:%d, desc:%s", resp.StatusCode(), string(resp.Body()))
	}

	return resp.Body(), nil
}

// new *fasthttp.Request
func (c *Client) newRequest(r api.IRequest) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	sign := c.newSignature(r)

	headers := map[string]string{
		fasthttp.HeaderContentType: "application/json;charset=utf-8",
		fasthttp.HeaderAccept:      "application/json",
		"OK-ACCESS-KEY":            c.Key,
		"OK-ACCESS-PASSPHRASE":     c.Passphrase,
		"OK-ACCESS-SIGN":           sign.Build(),
		"OK-ACCESS-TIMESTAMP":      sign.Timestamp,
	}
	if c.Simulated {
		headers["x-simulated-trading"] = "1"
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.SetMethod(sign.Method)
	req.SetRequestURI(c.Host + sign.Path)
	if sign.Body != "" {
		req.SetBodyString(sign.Body)
	}

	return req
}

// new *Signature
func (c *Client) newSignature(r api.IRequest) *Signature {
	var body []byte
	path := r.GetPath()

	if r.IsPost() {
		body, _ = json.Marshal(r.GetParam())
	} else if values, _ := query.Values(r.GetParam()); len(values) > 0 {
		path += "?" + values.Encode()
	}

	return &Signature{
		Key:    c.SecretKey,
		Method: r.GetMethod(),
		Path:   path,
		Body:   string(body),
	}
}
