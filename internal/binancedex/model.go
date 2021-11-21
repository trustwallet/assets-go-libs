package binancedex

type (
	Bep2Assets struct {
		ContractAddress  string `json:"contract_address,omitempty"`
		Name             string `json:"name"`
		OriginalSymbol   string `json:"original_symbol"`
		Owner            string `json:"owner"`
		Symbol           string `json:"symbol"`
		TotalSupply      string `json:"total_supply"`
		ContractDecimals int    `json:"contract_decimals,omitempty"`
		Mintable         bool   `json:"mintable"`
	}

	Bep8Assets struct {
		Name           string `json:"name"`
		OriginalSymbol string `json:"original_symbol"`
		Symbol         string `json:"symbol"`
		Owner          string `json:"owner"`
		TokenURI       string `json:"token_uri"`
		TotalSupply    string `json:"total_supply"`
		TokenType      int    `json:"token_type"`
		Mintable       bool   `json:"mintable"`
	}
)
