package path

import "fmt"

const githubUserContentURL = "https://raw.githubusercontent.com"

func GetAssetInfoGithubURL(repoOwner, repoName, branch, chain, tokenID string) string {
	return fmt.Sprintf("%s/%s/%s/%s/blockchains/%s/assets/%s/info.json",
		githubUserContentURL, repoOwner, repoName, branch, chain, tokenID)
}

func GetAssetLogoGithubURL(repoOwner, repoName, branch, chain, tokenID string) string {
	return fmt.Sprintf("%s/%s/%s/%s/blockchains/%s/assets/%s/logo.png",
		githubUserContentURL, repoOwner, repoName, branch, chain, tokenID)
}

func GetValidatorAssetLogoGithubURL(repoOwner, repoName, branch, chain, tokenID string) string {
	return fmt.Sprintf("%s/%s/%s/%s/blockchains/%s/validators/assets/%s/logo.png",
		githubUserContentURL, repoOwner, repoName, branch, chain, tokenID)
}

func GetValidatorListGithubURL(repoOwner, repoName, branch, chain string) string {
	return fmt.Sprintf("%s/%s/%s/%s/blockchains/%s/validators/list.json",
		githubUserContentURL, repoOwner, repoName, branch, chain)
}

func GetChainLogoURL(host, chain string) string {
	return fmt.Sprintf("%s/blockchains/%s/info/logo.png", host, chain)
}

func GetAssetLogoURL(host, chain, tokenID string) string {
	return fmt.Sprintf("%s/blockchains/%s/assets/%s/logo.png", host, chain, tokenID)
}
