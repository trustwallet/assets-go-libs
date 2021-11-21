package info

type (
	CoinModel struct {
		Name        *string  `json:"name,omitempty"`
		Website     *string  `json:"website,omitempty"`
		Description *string  `json:"description,omitempty"`
		Explorer    *string  `json:"explorer,omitempty"`
		Research    string   `json:"research,omitempty"`
		Symbol      *string  `json:"symbol,omitempty"`
		Type        *string  `json:"type,omitempty"`
		Decimals    *int     `json:"decimals,omitempty"`
		Status      *string  `json:"status,omitempty"`
		Tags        []string `json:"tags,omitempty"`
		Links       []Link   `json:"links,omitempty"`
	}

	Link struct {
		Name *string `json:"name,omitempty"`
		URL  *string `json:"url,omitempty"`
	}

	AssetModel struct {
		Name          *string  `json:"name,omitempty"`
		Symbol        *string  `json:"symbol,omitempty"`
		Type          *string  `json:"type,omitempty"`
		Decimals      *int     `json:"decimals,omitempty"`
		Description   *string  `json:"description,omitempty"`
		Website       *string  `json:"website,omitempty"`
		Explorer      *string  `json:"explorer,omitempty"`
		Status        *string  `json:"status,omitempty"`
		ID            *string  `json:"id,omitempty"`
		Research      string   `json:"research"`
		Links         []Link   `json:"links,omitempty"`
		ShortDesc     string   `json:"short_desc"`
		Audit         string   `json:"audit"`
		AuditReport   string   `json:"audit_report"`
		Tags          []string `json:"tags"`
		Code          string   `json:"code"`
		Ticker        string   `json:"ticker"`
		ExplorerEth   string   `json:"explorer-ETH"`
		Address       string   `json:"address"`
		Twitter       string   `json:"twitter"`
		CoinMarketcap string   `json:"coinmarketcap"`
		DataSource    string   `json:"data_source"`
	}
)

func (a *AssetModel) GetStatus() string {
	if a.Status == nil {
		return ""
	}

	return *a.Status
}
