package core

import "github.com/trustwallet/assets-go-libs/pkg/file"

type (
	Validator struct {
		Name     string
		FileType string

		Run func(f *file.AssetFile) error
	}

	Fixer struct {
		Name     string
		FileType string

		Run func(f *file.AssetFile) error
	}

	UpdaterAuto struct {
		Name string

		Run func() error
	}
)

type (
	TokenList struct {
		Name      string      `json:"name"`
		LogoURI   string      `json:"logoURI"`
		Timestamp string      `json:"timestamp"`
		Tokens    []TokenItem `json:"tokens"`
		Version   Version     `json:"version"`
	}

	TokenItem struct {
		Asset    string `json:"asset"`
		Type     string `json:"type"`
		Address  string `json:"address"`
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals uint   `json:"decimals"`
		LogoURI  string `json:"logoURI"`
		Pairs    []Pair `json:"pairs"`
	}

	Pair struct {
		Base     string `json:"base"`
		LotSize  string `json:"lotSize"`
		TickSize string `json:"tickSize"`
	}

	Version struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	}
)
