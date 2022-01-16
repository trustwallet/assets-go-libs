package tokenlist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/trustwallet/assets-go-libs/path"
	"github.com/trustwallet/assets-go-libs/validation"
	"github.com/trustwallet/assets-go-libs/validation/info"
	"github.com/trustwallet/go-primitives/asset"
	"github.com/trustwallet/go-primitives/coin"
	"github.com/trustwallet/go-primitives/types"
)

func ValidateTokenList(model Model, chain coin.Coin, tokenListPath string) error {
	compErr := validation.NewErrComposite()

	for _, token := range model.Tokens {
		err := validateTokenAddress(chain, token)
		if err != nil {
			compErr.Append(err)
		}

		err = validateChainOrAssetInfo(token, chain, tokenListPath)
		if err != nil {
			compErr.Append(err)
		}
	}

	err := validateTokenListPairs(model)
	if err != nil {
		compErr.Append(err)
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func validateChainOrAssetInfo(token Token, chain coin.Coin, tokenListPath string) error {
	var assetPath string

	if token.Type == types.Coin {
		assetPath = path.GetChainInfoPath(chain.Handle)
	} else {
		assetPath = path.GetAssetInfoPath(chain.Handle, token.Address)
	}

	infoFile, err := os.Open(assetPath)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	if _, err = buf.ReadFrom(infoFile); err != nil {
		return err
	}

	infoFile.Close()

	var infoAsset info.AssetModel
	err = json.Unmarshal(buf.Bytes(), &infoAsset)
	if err != nil {
		return err
	}

	if string(token.Type) != *infoAsset.Type {
		return fmt.Errorf("field type - '%s' differs from '%s' in %s",
			token.Type, *infoAsset.Type, assetPath)
	}

	if token.Symbol != *infoAsset.Symbol {
		return fmt.Errorf("field symbol - '%s' differs from '%s' in %s",
			token.Symbol, *infoAsset.Symbol, assetPath)
	}

	if token.Decimals != uint(*infoAsset.Decimals) {
		return fmt.Errorf("field decimals - '%d' differs from '%d' in %s",
			token.Decimals, *infoAsset.Decimals, assetPath)
	}

	if token.Name != *infoAsset.Name {
		return fmt.Errorf("field name - '%s' differs from '%s' in %s",
			token.Name, *infoAsset.Name, assetPath)
	}

	if infoAsset.GetStatus() != "active" {
		return fmt.Errorf("token '%s' is not active, remove it from %s", token.Address, tokenListPath)
	}

	return nil
}

func validateTokenListPairs(model Model) error {
	compErr := validation.NewErrComposite()

	tokensMap := make(map[string]struct{})
	for _, t := range model.Tokens {
		tokensMap[t.Asset] = struct{}{}
	}

	pairs := make(map[string]string)
	for _, t := range model.Tokens {
		for _, pair := range t.Pairs {
			pairs[pair.Base] = t.Address
		}
	}

	for pairToken, token := range pairs {
		if _, exists := tokensMap[pairToken]; !exists {
			compErr.Append(fmt.Errorf("token '%s' contains non-existing pair token '%s'", token, pairToken))
		}
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func validateTokenAddress(chain coin.Coin, token Token) error {
	if coin.IsEVM(chain.ID) {
		err := validation.ValidateETHForkAddress(chain, token.Address)
		if err != nil {
			return err
		}

		err = validateAssetID(chain, token.Asset)
		if err != nil {
			return err
		}

		for _, pair := range token.Pairs {
			err = validateAssetID(chain, pair.Base)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func validateAssetID(chain coin.Coin, id string) error {
	_, addr, err := asset.ParseID(id)
	if err != nil {
		return err
	}

	err = validation.ValidateETHForkAddress(chain, addr)
	if err != nil {
		return err
	}

	return nil
}
