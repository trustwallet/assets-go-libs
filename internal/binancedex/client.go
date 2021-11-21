package binancedex

import (
	"net/url"
	"strconv"

	"github.com/trustwallet/go-libs/client"
)

type Client struct {
	req client.Request
}

func InitBinanceDexClient(url string, errorHandler client.HttpErrorHandler) *Client {
	return &Client{
		req: client.InitClient(url, errorHandler),
	}
}

func (c *Client) GetBep2Assets(limit int) (assets []Bep2Assets, err error) {
	params := url.Values{"limit": {strconv.Itoa(limit)}}
	err = c.req.Get(&assets, "/api/v1/tokens", params)

	return assets, err
}

func (c *Client) GetBep8Assets(limit int) (assets []Bep8Assets, err error) {
	params := url.Values{"limit": {strconv.Itoa(limit)}}
	err = c.req.Get(&assets, "/api/v1/mini/tokens", params)

	return assets, err
}
