package info

import (
	"github.com/trustwallet/assets-go-libs/pkg/validation"
	"github.com/trustwallet/go-primitives/coin"
)

func ValidateCoin(model CoinModel, chain coin.Coin, addr string, allowedTags []string) error {
	if err := ValidateCoinRequiredKeys(model); err != nil {
		return err
	}

	// All fields validated for nil and can be safety used in func.
	compErr := validation.NewErrComposite()
	if err := ValidateCoinType(*model.Type); err != nil {
		compErr.Append(err)
	}

	if err := ValidateDecimals(*model.Decimals); err != nil {
		compErr.Append(err)
	}

	if err := ValidateStatus(*model.Status); err != nil {
		compErr.Append(err)
	}

	if err := ValidateCoinTags(model.Tags, allowedTags); err != nil {
		compErr.Append(err)
	}

	if err := ValidateDescription(*model.Description); err != nil {
		compErr.Append(err)
	}

	if err := ValidateDescriptionWebsite(*model.Description, *model.Website); err != nil {
		compErr.Append(err)
	}

	if err := ValidateCoinLinks(model.Links); err != nil {
		compErr.Append(err)
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}
