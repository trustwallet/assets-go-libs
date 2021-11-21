package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/assets-go-libs/internal/binancedex"
	"github.com/trustwallet/assets-go-libs/internal/config"
	"github.com/trustwallet/assets-go-libs/internal/processor"
	"github.com/trustwallet/assets-go-libs/internal/validators"
	"github.com/trustwallet/assets-go-libs/pkg/assetfs"
)

func main() {
	const defaultConfigPath = "configs/example.config.yaml"

	var confPath string
	flag.StringVar(&confPath, "c", defaultConfigPath, "config file")
	flag.Parse()

	conf, err := config.Load(confPath)
	if err != nil {
		log.WithError(err).Fatal("Failed to load app configuration")
	}

	logLevel, err := log.ParseLevel(conf.App.LogLevel)
	if err != nil {
		log.WithError(err).Fatal("Failed to parse log level")
	}
	log.SetLevel(logLevel)

	log.SetOutput(os.Stdin)

	log.WithField("conf", conf).Debug("App Config")

	binancedexClient := binancedex.InitBinanceDexClient(conf.ClientURLs.Binancedex, nil)
	var binanceAssetsSymbols []string

	bep8, err := binancedexClient.GetBep8Assets(1000)
	if err != nil {
		panic(err)
	}

	<-time.After(time.Second * 1) //TODO binance dex blocked too often requests (request timout)

	bep2, err := binancedexClient.GetBep2Assets(1000)
	if err != nil {
		panic(err)
	}

	for _, a := range bep8 {
		binanceAssetsSymbols = append(binanceAssetsSymbols, a.Symbol)
	}

	for _, a := range bep2 {
		binanceAssetsSymbols = append(binanceAssetsSymbols, a.Symbol)
	}

	paths, err := ReadLocalFileStructure()
	if err != nil {
		log.WithError(err).Fatal("Failed to load file structure")
	}

	fileStorage := assetfs.NewFileProvider()
	validatorsService := validators.NewService(
		conf.ValidatorsSettings,
		binanceAssetsSymbols,
		fileStorage,
	)

	//reporterService := reporter.NewReportService()

	assetfsProcessor := processor.NewService(conf.ValidatorsSettings, fileStorage, validatorsService)
	err = assetfsProcessor.RunSanityCheck(paths)
	if err != nil {
		log.WithError(err).Error()
	}
}

func ReadLocalFileStructure() ([]string, error) {
	var paths = []string{"./"}
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Contains("node_modules", path) {
				return nil
			}

			paths = append(paths, fmt.Sprintf("./%s", path))

			return nil
		})

	if err != nil {
		return nil, err
	}

	return paths, nil
}
