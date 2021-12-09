package dex

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

func (c *Client) GetMarketPairs(limit int) (pairs []MarketPair, err error) {
	params := url.Values{
		"limit": {strconv.Itoa(limit)},
	}
	err = c.req.Get(&pairs, "/v1/markets", params)

	return pairs, err
}

func (c *Client) GetTokensList(limit int) (tokens []Token, err error) {
	params := url.Values{
		"limit": {strconv.Itoa(limit)},
	}
	err = c.req.Get(&tokens, "/v1/tokens", params)

	return tokens, err
}
