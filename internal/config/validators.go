package config

type ImageFile struct {
	Size int `mapstructure:"size,omitempty"`
	MaxH int `mapstructure:"max_h,omitempty"`
	MaxW int `mapstructure:"max_w,omitempty"`
	MinH int `mapstructure:"min_h,omitempty"`
	MinW int `mapstructure:"min_w,omitempty"`
}

type AssetFolder struct {
	AllowedFiles []string `mapstructure:"allowed_files,omitempty"`
}

type RootFolder struct {
	AllowedFiles []string `mapstructure:"allowed_files,omitempty"`
	SkipFiles    []string `mapstructure:"skip_files,omitempty"`
}

type ChainFolder struct {
	AllowedFiles []string `mapstructure:"allowed_files,omitempty"`
}

type ChainInfoFolder struct {
	HasFiles []string `mapstructure:"has_files,omitempty"`
}

type DappsFolder struct {
	Ext string `mapstructure:"ext,omitempty"`
}

type CoinInfoFile struct {
	Tags []Tag `mapstructure:"tags,omitempty"`
}

type Tag struct {
	ID          string `mapstructure:"id,omitempty"`
	Name        string `mapstructure:"name,omitempty"`
	Description string `mapstructure:"description,omitempty"`
}
