package binancedex

type Bep2Asset struct {
	Asset       string `json:"asset,omitempty"`
	Name        string `json:"name,omitempty"`
	AssetImg    string `json:"assetImg,omitempty"`
	MappedAsset string `json:"mappedAsset,omitempty"`
	Decimals    int    `json:"decimals,omitempty"`
}

type Bep2Assets struct {
	AssetInfoList []Bep2Asset `json:"assetInfoList"`
}
