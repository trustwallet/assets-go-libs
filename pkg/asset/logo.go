package asset

import (
	"fmt"
	"strings"

	"github.com/trustwallet/go-primitives/types"
)

func GetAssetLogoPath(chain, tokenID string) string {
	return fmt.Sprintf("blockchains/%s/assets/%s/logo.png", chain, tokenID)
}

func GetAssetLogoURL(repoOwner, repoName, branch, chain, tokenID string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/blockchains/%s/assets/%s/logo.png",
		repoOwner, repoName, branch, chain, tokenID)
}

func GetTokenFromAssetLogoPath(path string) (tokenID, tokenType string, err error) {
	for _, t := range types.GetTokenTypes() {
		chain, err := types.GetChainFromAssetType(string(t))
		if err != nil {
			continue
		}

		prefix := fmt.Sprintf("blockchains/%s/assets/", chain.Handle)
		suffix := "/logo.png"

		if strings.HasPrefix(path, prefix) && strings.HasSuffix(path, suffix) {
			tokenType = string(t)
			tokenID = path[len(prefix):(len(path) - len(suffix))]

			break
		}
	}

	if tokenType == "TRC20" && strings.HasPrefix(tokenID, "10") {
		tokenType = "TRC10"
	}

	if tokenID == "" && tokenType == "" {
		return tokenID, tokenType, fmt.Errorf("couldn't get token id and token type")
	}

	return tokenID, tokenType, nil
}
