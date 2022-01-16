package tokenlist

import "github.com/trustwallet/go-primitives/types"

type (
	Model struct {
		Name      string  `json:"name"`
		LogoURI   string  `json:"logoURI"`
		Timestamp string  `json:"timestamp"`
		Tokens    []Token `json:"tokens"`
		Version   Version `json:"version"`
	}

	Token struct {
		Asset    string          `json:"asset"`
		Type     types.TokenType `json:"type"`
		Address  string          `json:"address"`
		Name     string          `json:"name"`
		Symbol   string          `json:"symbol"`
		Decimals uint            `json:"decimals"`
		LogoURI  string          `json:"logoURI"`
		Pairs    []Pair          `json:"pairs"`
	}

	Pair struct {
		Base     string `json:"base"`
		LotSize  string `json:"lotSize,omitempty"`
		TickSize string `json:"tickSize,omitempty"`
	}

	Version struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	}
)
