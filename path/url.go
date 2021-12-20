package path

import "fmt"

func GetAssetInfoGithubURL(repoOwner, repoName, branch, chain, tokenID string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/blockchains/%s/assets/%s/info.json",
		repoOwner, repoName, branch, chain, tokenID)
}

func GetAssetLogoGithubURL(repoOwner, repoName, branch, chain, tokenID string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/blockchains/%s/assets/%s/logo.png",
		repoOwner, repoName, branch, chain, tokenID)
}

func GetChainLogoURL(host, chain string) string {
	return fmt.Sprintf("%s/blockchains/%s/info/logo.png", host, chain)
}

func GetAssetLogoURL(host, chain, tokenID string) string {
	return fmt.Sprintf("%s/blockchains/%s/assets/%s/logo.png", host, chain, tokenID)
}
