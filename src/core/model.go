package core

import "github.com/trustwallet/assets-go-libs/pkg/file"

type (
	Validator struct {
		Name     string
		FileType string

		Run func(f *file.AssetFile) error
	}

	Fixer struct {
		Validator
	}

	UpdaterAuto struct {
		Name string

		Run func() error
	}
)

type (
	TokenItem struct {
		Asset    string
		Type     string
		Address  string
		Name     string
		Symbol   string
		LogoURI  string
		Decimals uint
		Pairs    []Pair
	}

	Pair struct {
		Base     string
		LotSize  int64
		TickSize int64
	}
)
