package entity

import (
	"net/url"
)

// Client url builder
type Client struct {
	URL url.URL
}

// NewClient create new Client
func NewClient(scheme, host, apiKey string) *Client {
	var u url.URL
	u.Scheme = scheme
	u.Host = host
	u.Path = "/api"
	q := u.Query()
	q.Add("apikey", apiKey)
	u.RawQuery = q.Encode()
	return &Client{
		URL: u,
	}
}

//AddParams adds params to url
func (c *Client) AddParams(params map[string]string) {
	q := c.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	c.URL.RawQuery = q.Encode()
}
