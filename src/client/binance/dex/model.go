package dex

type (
	MarketPair struct {
		BaseAssetSymbol  string `json:"base_asset_symbol"`
		LotSize          string `json:"lot_size"`
		QuoteAssetSymbol string `json:"quote_asset_symbol"`
		TickSize         string `json:"tick_size"`
	}

	Token struct {
		Symbol         string `json:"symbol"`
		Name           string `json:"name"`
		OriginalSymbol string `json:"original_symbol"`
	}
)
