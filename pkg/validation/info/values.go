package info

import (
	"fmt"
	"strings"

	"github.com/trustwallet/go-primitives/coin"
)

var requiredCoinFields = []string{"name", "type", "symbol", "decimals", "description", "website", "explorer", "status"}
var requiredAssetFields = []string{"name", "type", "symbol", "decimals", "description", "website", "explorer", "status", "id"}
var allowedStatusValues = []string{
	"active",
	"spam",
	"abandoned",
}
var allowedLinkKeys = map[string]string{"github": "https://github.com/",
	"whitepaper":    "",
	"twitter":       "https://twitter.com/",
	"telegram":      "https://t.me/",
	"telegram_news": "https://t.me/", // read-only announcement channel
	"medium":        "",              // url contains 'medium.com'
	"discord":       "https://discord.com/",
	"reddit":        "https://reddit.com/",
	"facebook":      "https://facebook.com/",
	"youtube":       "https://youtube.com/",
	"coinmarketcap": "https://coinmarketcap.com/",
	"coingecko":     "https://coingecko.com/",
	"blog":          "", // blog, other than medium
	"forum":         "", // community site
	"docs":          "",
	"source_code":   "", // other than github
}

func explorerUrlAlternatives(chain string, name string) []string {
	var altUrls []string

	if name != "" {
		NameNorm := strings.Replace(strings.Replace(strings.Replace(strings.ToLower(name), " ", "", -1), ")", "", -1), "(", "", -1)
		if strings.ToLower(chain) == coin.Coins[coin.ETHEREUM].Name {
			altUrls = append(altUrls, fmt.Sprintf("https://etherscan.io/token/%s", NameNorm))
		}

		altUrls = append(altUrls, fmt.Sprintf("https://explorer.%s.io", NameNorm))
		altUrls = append(altUrls, fmt.Sprintf("https://scan.%s.io", NameNorm))
	}

	return altUrls
}

func linkNameAllowed(str string) bool {
	if _, exists := allowedLinkKeys[str]; !exists {
		return false
	}

	return true
}

func supportedLinkNames() []string {
	var names []string
	for k, _ := range allowedLinkKeys {
		names = append(names, k)
	}

	return names
}

func supportedLinkValues() []string {
	var values []string
	for _, v := range allowedLinkKeys {
		values = append(values, v)
	}

	return values
}
