package path_test

import (
	"testing"

	"github.com/trustwallet/assets-go-libs/path"
)

func Test_GetTokenFromAssetLogoPath(t *testing.T) {
	tests := []struct {
		name            string
		path            string
		resulTokenID    string
		resultTokenType string
	}{
		{
			name:            "Tron (TRC20)",
			path:            "blockchains/tron/assets/TXBcx59eDVndV5upFQnTR2xdvqFd5reXET/logo.png",
			resulTokenID:    "TXBcx59eDVndV5upFQnTR2xdvqFd5reXET",
			resultTokenType: "TRC20",
		},
		{
			name:            "Ethereum (ERC20)",
			path:            "blockchains/ethereum/assets/0x0a2D9370cF74Da3FD3dF5d764e394Ca8205C50B6/logo.png",
			resulTokenID:    "0x0a2D9370cF74Da3FD3dF5d764e394Ca8205C50B6",
			resultTokenType: "ERC20",
		},
		{
			name:            "Binance (BEP2)",
			path:            "blockchains/binance/assets/AAVE-8FA/logo.png",
			resulTokenID:    "AAVE-8FA",
			resultTokenType: "BEP2",
		},
		{
			name:            "Smartchain (BEP20)",
			path:            "blockchains/smartchain/assets/0x0b3f42481C228F70756DbFA0309d3ddC2a5e0F6a/logo.png",
			resulTokenID:    "0x0b3f42481C228F70756DbFA0309d3ddC2a5e0F6a",
			resultTokenType: "BEP20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenID, tokenType := path.GetTokenFromAssetLogoPath(tt.path)
			if tokenID != tt.resulTokenID || tokenType != tt.resultTokenType {
				t.Errorf("got tokenID=%s tokenType=%s, want tokenID=%s tokenType=%s",
					tokenID, tokenType, tt.resulTokenID, tt.resultTokenType)
			}
		})
	}
}
