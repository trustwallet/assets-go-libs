package path

import (
	"fmt"
	"strings"

	"github.com/trustwallet/go-primitives/types"
)

type TokenListType int

const (
	TokenlistDefault = iota
	TokenlistExtended
)

func GetTokenListPath(chain string, tokenlistType TokenListType) string {
	switch tokenlistType {
	case TokenlistDefault:
		return fmt.Sprintf("blockchains/%s/tokenlist.json", chain)
	case TokenlistExtended:
		return fmt.Sprintf("blockchains/%s/tokenlist-extended.json", chain)
	}

	return ""
}

func GetAssetPath(chain, tokenID string) string {
	return fmt.Sprintf("blockchains/%s/assets/%s", chain, tokenID)
}

func GetAssetInfoPath(chain, tokenID string) string {
	return fmt.Sprintf("blockchains/%s/assets/%s/info.json", chain, tokenID)
}

func GetAssetLogoPath(chain, tokenID string) string {
	return fmt.Sprintf("blockchains/%s/assets/%s/logo.png", chain, tokenID)
}

func GetChainInfoPath(chain string) string {
	return fmt.Sprintf("blockchains/%s/info/info.json", chain)
}

func GetValidatorAssetsPath(chain string) string {
	return fmt.Sprintf("blockchains/%s/validators/assets", chain)
}

func GetTokenFromAssetLogoPath(path string) (tokenID, tokenType string) {
	for _, t := range types.GetTokenTypes() {
		chain, err := types.GetChainFromAssetType(string(t))
		if err != nil {
			continue
		}

		prefix := fmt.Sprintf("blockchains/%s/assets/", chain.Handle)
		suffix := "/logo.png"

		if strings.HasPrefix(path, prefix) && strings.HasSuffix(path, suffix) {
			tokenID = path[len(prefix):(len(path) - len(suffix))]

			var ok bool
			tokenType, ok = types.GetTokenType(chain.ID, tokenID)
			if !ok {
				tokenType = string(t)
			}

			break
		}
	}

	return tokenID, tokenType
}
