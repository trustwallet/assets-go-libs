package explorer

import (
	"net/url"
	"strconv"

	"github.com/trustwallet/go-libs/client"
)

type Client struct {
	req client.Request
}

func NewClient(url string, errorHandler client.HttpErrorHandler) *Client {
	return &Client{req: client.InitClient(url, errorHandler)}
}

func (c *Client) GetBep2Assets(page, rows int) (assets Bep2Assets, err error) {
	params := url.Values{
		"page": {strconv.Itoa(page)},
		"rows": {strconv.Itoa(rows)},
	}
	err = c.req.Get(&assets, "/api/v1/assets", params)

	return assets, err
}