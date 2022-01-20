package validation

import (
	"fmt"
	"regexp"
	"strings"

	str "github.com/trustwallet/assets-go-libs/strings"
	"github.com/trustwallet/go-primitives/address"
	"github.com/trustwallet/go-primitives/coin"
)

var regexTRC10 = regexp.MustCompile(`^\d+$`)

func ValidateAssetAddress(chain coin.Coin, address string) error {
	switch {
	case coin.IsEVM(chain.ID):
		return ValidateETHForkAddress(chain, address)
	case chain.ID == coin.TRON:
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

func ValidateTezosAddress(addr string) error {
	if !strings.HasPrefix(addr, "tz") {
		return fmt.Errorf("%w: shoud be valid tezos address", ErrInvalidAddress)
	}

	return nil
}

func ValidateTronAddress(addr string) error {
	trc20 := len(addr) == 34 && strings.HasPrefix(addr, "T") && !str.IsLowerCase(addr) && !str.IsUpperCase(addr)
	trc10 := regexTRC10.MatchString(addr)

	if !trc10 && !trc20 {
		return fmt.Errorf("%w: should be valid tron address", ErrInvalidAddress)
	}

	return nil
}

func ValidateWavesAddress(addr string) error {
	condition := strings.HasPrefix(addr, "3P") && len(addr) == 35 && !str.IsLowerCase(addr) && !str.IsUpperCase(addr)
	if !condition {
		return fmt.Errorf("%w: should be valid waves address", ErrInvalidAddress)
	}

	return nil
}

func ValidateETHForkAddress(chain coin.Coin, addr string) error {
	checksum, err := address.EIP55Checksum(addr)
	if err != nil {
		return fmt.Errorf("failed to get address checksum: %w", err)
	}

	if chain.ID == coin.WANCHAIN {
		checksum = strings.ReplaceAll(str.ReverseCase(checksum), "X", "x")
	}

	if checksum != addr {
		return fmt.Errorf("%w: expect asset %s in checksum: %s", ErrInvalidAddress, addr, checksum)
	}

	return nil
}

func ValidateAddress(address, prefix string, length int) error {
	if !strings.HasPrefix(address, prefix) {
		return fmt.Errorf("%w: %s should has prefix %s", ErrInvalidFileNameCase, address, prefix)
	}

	if len(address) != length {
		return fmt.Errorf("%w: %s should be %d length", ErrInvalidFileNameLength, address, length)
	}

	if !str.IsLowerCase(address) {
		return fmt.Errorf("%w: %s should be lowercase", ErrInvalidFileNameCase, address)
	}

	return nil
}

func IsEthereumAddress(addr string) bool {
	if len(addr) == 40 || len(addr) == 42 && strings.HasPrefix(addr, "0x") {
		return true
	}

	return false
}
