package config

import (
	"path/filepath"

	"github.com/trustwallet/go-libs/config/viper"
)

type Config struct {
	App                App                `mapstructure:"app"`
	ValidatorsSettings ValidatorsSettings `mapstructure:"validators_settings"`
	ClientURLs         ClientsURLs        `mapstructure:"client_urls"`
}

type App struct {
	Mode     string `mapstructure:"mode"`
	LogLevel string `mapstructure:"log_level"`
}

type ClientsURLs struct {
	Binancedex string `mapstructure:"binancedex"`
}

type ValidatorsSettings struct {
	CoinInfoFile    CoinInfoFile    `mapstructure:"coin_info_file"`
	ImageFile       ImageFile       `mapstructure:"image_file"`
	AssetFolder     AssetFolder     `mapstructure:"asset_folder"`
	ChainFolder     ChainFolder     `mapstructure:"chain_folder"`
	RootFolder      RootFolder      `mapstructure:"root_folder"`
	ChainInfoFolder ChainInfoFolder `mapstructure:"chain_info_folder"`
	DaapsFolder     DaapsFolder     `mapstructure:"dapps_folder"`
}

type AppMode string

func Load(confPath string) (*Config, error) {
	confPath, err := filepath.Abs(confPath)
	if err != nil {
		return nil, err
	}

	var config Config
	viper.Load(confPath, &config)

	return &config, nil
}

func (mode AppMode) IsDev() bool {
	return mode == "dev"
}
