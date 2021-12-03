package external

import (
	"fmt"
	"strconv"

	"github.com/trustwallet/assets-go-libs/pkg"
)

type TokenInfoERC20 struct {
	Symbol       string `json:"symbol"`
	Decimals     string `json:"decimals"`
	HoldersCount int    `json:"holdersCount"`
}

func GetTokenInfoForERC20(tokenID string) (*TokenInfo, error) {
	url := fmt.Sprintf("https://api.ethplorer.io/getTokenInfo/%s?apiKey=freekey", tokenID)

	var result TokenInfoERC20
	err := pkg.GetHTTPResponse(url, &result)
	if err != nil {
		return nil, err
	}

	decimals, err := strconv.Atoi(result.Decimals)
	if err != nil {
		return nil, err
	}

	return &TokenInfo{
		Symbol:       result.Symbol,
		Decimals:     decimals,
		HoldersCount: result.HoldersCount,
	}, nil
}
