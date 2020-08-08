package robot

import (
	"encoding/json"
	"io/ioutil"
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

type Result struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// Send send message
func (c *Client) Send(m *Msg) (*Result, error) {
	mr, err := m.Reader()
	if err != nil {
		return nil, err
	}
	resp, err := c.c.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+c.key, "application/json", mr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := &Result{}
	err = json.Unmarshal(respData, r)
	return r, err
}
