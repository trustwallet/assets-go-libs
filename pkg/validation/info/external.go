package info

import (
	"fmt"
	"strings"

	"github.com/trustwallet/assets-go-libs/pkg"
)

type ExternalTokenInfo struct {
	Symbol       string `json:"symbol"`
	Decimals     int    `json:"decimals"`
	HoldersCount int    `json:"holdersCount"`
}

func GetExternalTokenInfo(tokenID, tokentType string) (*ExternalTokenInfo, error) {
	switch strings.ToLower(tokentType) {
	case "erc20":
		return GetTokenInfoForERC20(tokenID)
	case "bep20":
		return GetTokenInfoForBEP20(tokenID)
	}

	return nil, nil
}

func GetTokenInfoForERC20(tokenID string) (*ExternalTokenInfo, error) {
	url := fmt.Sprintf("https://api.ethplorer.io/getTokenInfo/%s?apiKey=freekey", tokenID)

	var result ExternalTokenInfo
	err := pkg.GetHTTPResponse(url, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// TODO: Implement it.
func GetTokenInfoForBEP20(tokenID string) (*ExternalTokenInfo, error) {
	return nil, nil
}
