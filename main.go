package main

import (
	"github.com/Hertucktor/archive-importer/config"
	"github.com/Hertucktor/archive-importer/importer"
	"github.com/Hertucktor/archive-importer/utils"
)

func main() {
	var page = 1
	var delimiter = 1
	configFile := "config.yml"
	logger := utils.InitializeSugarLogger("")
	conf, err := config.GetConfig(configFile, logger)
	if err != nil {
		logger.Fatal(err)
	}

	if err := importer.ImportCardsIntoDatabase(conf, page, delimiter, logger); err != nil {
		logger.Fatal(err)
	}

}
