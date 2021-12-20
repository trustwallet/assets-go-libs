package external

import (
	"strings"
)

type TokenInfo struct {
	Symbol       string
	Decimals     int
	HoldersCount int
}

func GetTokenInfo(tokenID, tokentType string) (*TokenInfo, error) {
	switch strings.ToLower(tokentType) {
	case "erc20":
		return GetTokenInfoForERC20(tokenID)
	case "bep20":
		return GetTokenInfoForBEP20(tokenID)
	}

	return nil, nil
}
