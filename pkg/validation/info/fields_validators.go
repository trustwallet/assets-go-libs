package info

import (
	"fmt"
	"strings"

	"github.com/trustwallet/assets-go-libs/pkg"
	"github.com/trustwallet/go-primitives/coin"
	"github.com/trustwallet/go-primitives/types"
)

// Here is list of function for validate info.CoinModel and info.AssetModel structs.

// AssetModel info specific validators.

func ValidateAssetRequiredKeys(a AssetModel) error {
	var fields []string
	if a.Name != nil && !isEmpty(*a.Name) {
		fields = append(fields, "name")
	}
	if a.Symbol != nil && !isEmpty(*a.Symbol) {
		fields = append(fields, "symbol")
	}
	if a.Type != nil && !isEmpty(*a.Type) {
		fields = append(fields, "type")
	}
	if a.Decimals != nil {
		fields = append(fields, "decimals")
	}
	if a.Description != nil && !isEmpty(*a.Description) {
		fields = append(fields, "description")
	}
	if a.Website != nil {
		fields = append(fields, "website")
	}
	if a.Explorer != nil && !isEmpty(*a.Explorer) {
		fields = append(fields, "explorer")
	}
	if a.Status != nil && !isEmpty(*a.Status) {
		fields = append(fields, "status")
	}
	if a.ID != nil && !isEmpty(*a.ID) {
		fields = append(fields, "id")
	}

	if len(fields) != len(requiredAssetFields) {
		return fmt.Errorf("missing or empty required fields\n-%s",
			strings.Join(difference(requiredAssetFields, fields), "\n"))
	}

	return nil
}

func ValidateAssetType(type_ string, chain coin.Coin) error {
	chainFromType, err := types.GetChainFromAssetType(type_)
	if err != nil {
		return fmt.Errorf("invalid type field: %w", err)
	}

	if chainFromType != chain {
		return fmt.Errorf("invalid value for field type")
	}

	if strings.ToUpper(type_) != type_ {
		return fmt.Errorf("invalid value for type filed, should be ALLCAPS")
	}

	return nil
}

func ValidateAssetID(id string, address string) error {
	if id != address {
		if !strings.EqualFold(id, address) {
			return fmt.Errorf("invalid id field")
		}

		return fmt.Errorf("invalid casing for id field")
	}

	return nil
}

func ValidateAssetDecimalsAccordingType(type_ string, decimals int) error {
	if type_ == "BEP2" && decimals != 8 {
		return fmt.Errorf("invalid decimals field, BEP2 tokens have 8 decimals")
	}

	return nil
}

// CoinModel info specific validators.

func ValidateCoinRequiredKeys(c CoinModel) error {
	var fields []string
	if c.Name != nil && !isEmpty(*c.Name) {
		fields = append(fields, "name")
	}
	if c.Symbol != nil && !isEmpty(*c.Symbol) {
		fields = append(fields, "symbol")
	}
	if c.Type != nil && !isEmpty(*c.Type) {
		fields = append(fields, "type")
	}
	if c.Decimals != nil {
		fields = append(fields, "decimals")
	}
	if c.Description != nil && !isEmpty(*c.Description) {
		fields = append(fields, "description")
	}
	if c.Website != nil && !isEmpty(*c.Website) {
		fields = append(fields, "website")
	}
	if c.Explorer != nil && !isEmpty(*c.Explorer) {
		fields = append(fields, "explorer")
	}
	if c.Status != nil && !isEmpty(*c.Status) {
		fields = append(fields, "status")
	}

	if len(fields) != len(requiredCoinFields) {
		return fmt.Errorf("missing or empty required fields\n-%s",
			strings.Join(difference(requiredCoinFields, fields), "\n"))
	}

	return nil
}

func ValidateCoinLinks(links []Link) error {
	if len(links) < 1 {
		return nil
	}

	for _, l := range links {
		if l.Name == nil || l.URL == nil {
			return fmt.Errorf("missing required fields links.url and links.name")
		}

		if !linkNameAllowed(*l.Name) {
			return fmt.Errorf("invalid value for links.name filed, allowed only - %s",
				strings.Join(supportedLinkNames(), ", "))
		}

		prefix := allowedLinkKeys[*l.Name]
		if prefix != "" {
			if !strings.HasPrefix(*l.URL, prefix) {
				return fmt.Errorf("invalid value for links.url field, allowed only with prefixes - %s",
					strings.Join(supportedLinkValues(), ", "))
			}
		}

		if !strings.HasPrefix(*l.URL, "https://") {
			return fmt.Errorf("invalid value for links.url field, allowed only with https:// prefix")
		}

		if *l.Name == "medium" {
			if !strings.Contains(*l.URL, "medium.com") {
				return fmt.Errorf("invalid value for links.url field, should contain medium.com")
			}
		}
	}

	return nil
}

func ValidateCoinType(type_ string) error {
	if type_ != "coin" {
		return fmt.Errorf("invalid value for coin field, allowed only \"coin\"")
	}

	return nil
}

func ValidateCoinTags(tags []string, allowedTags []string) error {
	for _, t := range tags {
		if !pkg.Contains(t, allowedTags) {
			return fmt.Errorf("invalid value for tags field, tag %s - not allowed", t)
		}
	}

	return nil
}

// Both infos can be validated by this validators.

func ValidateDecimals(decimals int) error {
	if decimals > 30 || decimals < 0 {
		return fmt.Errorf("invalid value for decimals field")
	}

	return nil
}

func ValidateStatus(status string) error {
	for _, f := range allowedStatusValues {
		if f == status {
			return nil
		}
	}

	return fmt.Errorf("invalid value for status field")
}

func ValidateDescription(description string) error {
	if description == "" {
		return fmt.Errorf("invalid value for description field, for empty desciption use \"-\"")
	}

	if len(description) > 600 {
		return fmt.Errorf("invalid length for description field")
	}

	return nil
}

func ValidateDescriptionWebsite(description, website string) error {
	if description != "-" && website == "" {
		return fmt.Errorf("missing value for one of required fields - website")
	}

	return nil
}

func ValidateExplorer(explorer, name string, chain coin.Coin, addr string) error {
	explorerExpected, err := coin.GetCoinExploreURL(chain, addr)
	if err != nil {
		explorerExpected = ""
	}

	explorerActual := explorer

	if !strings.EqualFold(explorerActual, explorerExpected) {
		explorerAlt := explorerUrlAlternatives(chain.Handle, name)
		if len(explorerAlt) == 0 {
			return nil
		}

		var matchCount = 0

		for _, e := range explorerAlt {
			if strings.EqualFold(e, explorerActual) {
				matchCount++
			}
		}

		if matchCount == 0 {
			return fmt.Errorf("invalid value for explorer field, %s insted of %s", explorerActual, explorerExpected)
		}
	}

	return nil
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}

	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}

	return diff
}

func isEmpty(field string) bool {
	return field == ""
}
