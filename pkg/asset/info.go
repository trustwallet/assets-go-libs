package asset

import "fmt"

func GetAssetInfoPath(chain, tokenID string) string {
	return fmt.Sprintf("blockchains/%s/assets/%s/info.json", chain, tokenID)
}

func GetAssetInfoURL(repoOwner, repoName, branch, chain, tokenID string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/blockchains/%s/assets/%s/info.json",
		repoOwner, repoName, branch, chain, tokenID)
}
