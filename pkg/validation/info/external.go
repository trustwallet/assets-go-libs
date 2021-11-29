package info

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/trustwallet/assets-go-libs/pkg"
)

type ExternalTokenInfo struct {
	Symbol       string
	Decimals     int
	HoldersCount int
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

type ExternalTokenInfoERC20 struct {
	Symbol       string `json:"symbol"`
	Decimals     string `json:"decimals"`
	HoldersCount int    `json:"holdersCount"`
}

func GetTokenInfoForERC20(tokenID string) (*ExternalTokenInfo, error) {
	url := fmt.Sprintf("https://api.ethplorer.io/getTokenInfo/%s?apiKey=freekey", tokenID)

	var result ExternalTokenInfoERC20
	err := pkg.GetHTTPResponse(url, &result)
	if err != nil {
		return nil, err
	}

	decimals, err := strconv.Atoi(result.Decimals)
	if err != nil {
		return nil, err
	}

	return &ExternalTokenInfo{
		Symbol:       result.Symbol,
		Decimals:     decimals,
		HoldersCount: result.HoldersCount,
	}, nil
}

// TODO: Implement it.
func GetTokenInfoForBEP20(tokenID string) (*ExternalTokenInfo, error) {
	return nil, nil
}
