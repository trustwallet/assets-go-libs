package assetsmanager

import (
	"context"
	"time"

	"github.com/trustwallet/go-libs/client"
)

type Client struct {
	req client.Request
}

func InitClient(url string, errorHandler client.HttpErrorHandler) Client {
	return Client{
		req: client.InitJSONClient(url, errorHandler),
	}
}

func (c *Client) ValidateAssetInfo(req *AssetValidationReq) (result AssetValidationResp, err error) {
	request := client.NewReqBuilder().
		Method("POST").
		PathStatic("/api/v1/validate/asset_info").
		Body(req).
		WriteTo(&result).Build()

	_, err = c.req.Execute(context.Background(), request)

	return result, err
}

func (c *Client) GetTagValues() (result TagValuesResp, err error) {
	err = c.req.GetWithCache(&result, "/api/v1/values/tags", nil, time.Hour)

	return result, err
}
