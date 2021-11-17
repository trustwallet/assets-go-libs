package binancedex

import (
	"net/url"
	"strconv"

	"github.com/trustwallet/go-libs/client"
)

func InitBinanceDexClient(url string, errorHandler client.HttpErrorHandler) *Client {
	req := client.InitClient(url, errorHandler)

	return &Client{
		req: req,
	}
}

type Client struct {
	req client.Request
}

func (c *Client) GetBep2Assets(limit int) (assets []*Bep2Assets, err error) {
	params := url.Values{"limit": {strconv.Itoa(limit)}}
	err = c.req.Get(&assets, "/api/v1/tokens", params)

	return assets, err
}

func (c *Client) GetBep8Assets(limit int) (assets []*Bep8Assets, err error) {
	params := url.Values{"limit": {strconv.Itoa(limit)}}
	err = c.req.Get(&assets, "/api/v1/mini/tokens", params)

	return assets, err
}
