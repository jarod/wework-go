package robot

import (
	"net/http"
)

// Client wework robot client api
// https://work.weixin.qq.com/api/doc/90000/90136/91770
type Client struct {
	key string
	c   *http.Client
}

// New create client instance with given api key
func New(key string) *Client {
	return &Client{
		key: key,
		c:   &http.Client{},
	}
}

// Send send message
func (c *Client) Send(m *Msg) error {
	mr, err := m.Reader()
	if err != nil {
		return err
	}
	_, err = c.c.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+c.key, "application/json", mr)
	if err != nil {
		return err
	}
	return nil
}
