package validation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/trustwallet/assets-go-libs/pkg"
	"github.com/trustwallet/go-primitives/address"
	"github.com/trustwallet/go-primitives/coin"
)

var regexTRC10 = regexp.MustCompile(`^\d+$`)

func ValidateAssetAddress(chain coin.Coin, address string) error {
	switch chain.ID {
	case
		coin.ETHEREUM,
		coin.CLASSIC,
		coin.POA,
		coin.TOMOCHAIN,
		coin.GOCHAIN,
		coin.WANCHAIN,
		coin.THUNDERTOKEN,
		coin.SMARTCHAIN,
		coin.POLYGON,
		coin.OPTIMISM,
		coin.XDAI,
		coin.AVALANCHEC,
		coin.ARBITRUM,
		coin.FANTOM:
		return ValidateETHForkAddress(chain, address)
	case coin.TRON:
		return ValidateTronAddress(address)
	}

	return nil
}

func ValidateValidatorsAddress(chain coin.Coin, address string) error {
	switch chain.ID {
	case coin.TEZOS:
		return ValidateTezosAddress(address)
	case coin.TRON:
		return ValidateTronAddress(address)
	case coin.WAVES:
		return ValidateWavesAddress(address)
	case coin.COSMOS:
		return ValidateAddress(address, "cosmosvaloper1", 52)
	case coin.TERRA:
		return ValidateAddress(address, "terravaloper1", 51)
	case coin.KAVA:
		return ValidateAddress(address, "kavavaloper1", 50)
	}

	return nil
}

func ValidateWavesAddress(addr string) error {
	condition := strings.HasPrefix(addr, "3P") && len(addr) == 35 && !pkg.IsLowerCase(addr) && !pkg.IsUpperCase(addr)
	if !condition {
		return fmt.Errorf("%w, %s - should be Waves address", ErrInvalidAddress, addr)
	}

	return nil
}

func ValidateTezosAddress(addr string) error {
	if !strings.HasPrefix(addr, "tz") {
		return fmt.Errorf("%w, shoud be valid tezos address", ErrInvalidAddress)
	}

	return nil
}

func ValidateTronAddress(addr string) error {
	trc20 := len(addr) == 34 && strings.HasPrefix(addr, "T") && !pkg.IsLowerCase(addr) && !pkg.IsUpperCase(addr)
	trc10 := regexTRC10.MatchString(addr)

	if !trc10 && !trc20 {
		return fmt.Errorf("%w: should be valid tron address", ErrInvalidAddress)
	}

	return nil
}

func ValidateETHForkAddress(chain coin.Coin, addr string) error {
	checksum, err := address.EIP55Checksum(addr)
	if err != nil {
		return err
	}

	if chain.ID == coin.WANCHAIN {
		checksum = strings.ReplaceAll(pkg.ReverseCase(checksum), "X", "x")
	}

	if checksum != addr {
		return fmt.Errorf("expect asset %s in checksum: %s", addr, checksum)
	}

	return nil
}

func ValidateAddress(address string, prefix string, length int) error {
	if !strings.HasPrefix(address, prefix) {
		return fmt.Errorf("%w: %s should has prefix %s", ErrInvalidFileCase, address, prefix)
	}

	if len(address) != length {
		return fmt.Errorf("%w: %s should be %d length", ErrInvalidFileNameLength, address, length)
	}

	if !pkg.IsLowerCase(address) {
		return fmt.Errorf("%w: %s should be lowercase", ErrInvalidFileCase, address)
	}

	return nil
}
