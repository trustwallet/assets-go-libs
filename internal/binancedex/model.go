package binancedex

type Bep2Assets struct {
	ContractAddress  string `json:"contract_address,omitempty"`
	ContractDecimals int    `json:"contract_decimals,omitempty"`
	Mintable         bool   `json:"mintable"`
	Name             string `json:"name"`
	OriginalSymbol   string `json:"original_symbol"`
	Owner            string `json:"owner"`
	Symbol           string `json:"symbol"`
	TotalSupply      string `json:"total_supply"`
}

type Bep8Assets struct {
	Name           string `json:"name"`
	OriginalSymbol string `json:"original_symbol"`
	Symbol         string `json:"symbol"`
	Owner          string `json:"owner"`
	TokenURI       string `json:"token_uri"`
	TokenType      int    `json:"token_type"`
	TotalSupply    string `json:"total_supply"`
	Mintable       bool   `json:"mintable"`
}
