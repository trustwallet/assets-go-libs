package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/trustwallet/assets-go-libs/internal/config"
	"github.com/trustwallet/assets-go-libs/internal/processor"
	"github.com/trustwallet/assets-go-libs/internal/validator"
	"github.com/trustwallet/assets-go-libs/pkg/file"
)

var (
	configPath, root string
)

func main() {
	setup()

	paths, err := file.ReadLocalFileStructure(root, config.Default.ValidatorsSettings.RootFolder.SkipFiles)
	if err != nil {
		log.WithError(err).Fatal("failed to load file structure")
	}

	fileStorage := file.NewFileProvider()

	validatorsService, err := validator.NewService(fileStorage)
	if err != nil {
		log.WithError(err).Fatal("failed to init validator service")
	}

	// TODO
	// reporterService := reporter.NewReportService()

	assetfsProcessor := processor.NewService(fileStorage, validatorsService)
	err = assetfsProcessor.RunSanityCheck(paths)
	if err != nil {
		log.WithError(err).Error()
	}
}

func setup() {
	flag.StringVar(&configPath, "c", "", "path to config file")
	flag.StringVar(&root, "r", "./", "path to the root of the dir")
	flag.Parse()

	if err := config.SetConfig(configPath); err != nil {
		log.WithError(err).Fatal("failed to set config")
	}

	logLevel, err := log.ParseLevel(config.Default.App.LogLevel)
	if err != nil {
		log.WithError(err).Fatal("failed to parse log level")
	}

	log.SetLevel(logLevel)
	log.SetOutput(os.Stdin)

	log.WithField("conf", config.Default).Debug("App Config")
}
