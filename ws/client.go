package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	EndpointPublic           = "wss://ws.okx.com:8443/ws/v5/public"
	EndpointPrivate          = "wss://ws.okx.com:8443/ws/v5/private"
	EndpointPublicSimulated  = "wss://wspap.okx.com:8443/ws/v5/public?brokerId=9999"
	EndpointPrivateSimulated = "wss://wspap.okx.com:8443/ws/v5/private?brokerId=9999"
)

var (
	DefaultClientPublic           = NewClient(EndpointPublic)
	DefaultClientPrivate          = NewClient(EndpointPrivate)
	DefaultClientPublicSimulated  = NewClient(EndpointPublicSimulated)
	DefaultClientPrivateSimulated = NewClient(EndpointPrivateSimulated)
)

type Client struct {
	Endpoint string
	Dialer   *websocket.Dialer
}

// new Client
func NewClient(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
	}
}

// subscribe
func (c *Client) Subscribe(subscribe *Subscribe) error {
	conn, _, err := c.dial()
	if err != nil {
		return err
	}
	if err := c.messageSubscribe(conn, subscribe); err != nil {
		return err
	}

	go c.messageLoop(conn, subscribe)

	return nil
}

// message subscribe
func (c *Client) messageSubscribe(conn *websocket.Conn, subscribe *Subscribe) error {
	if err := conn.WriteJSON(subscribe.Request); err != nil {
		return err
	}
	if err := conn.ReadJSON(&subscribe.Response); err != nil {
		return err
	}
	return subscribe.Response.Error()
}

// loop websocket message
func (c *Client) messageLoop(conn *websocket.Conn, subscribe *Subscribe) {
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			subscribe.HandlerError(err)
			return
		}
		subscribe.Handler(message)
	}
}

// dial endpoint
func (c *Client) dial() (*websocket.Conn, *http.Response, error) {
	if c.Dialer == nil {
		c.Dialer = websocket.DefaultDialer
	}
	return c.Dialer.Dial(c.Endpoint, nil)
}
