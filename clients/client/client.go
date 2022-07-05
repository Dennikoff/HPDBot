package client

import "net/http"

type Client struct {
	host     string
	baseBath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		baseBath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Update() {

}

func (c *Client) SendMessage() {

}
