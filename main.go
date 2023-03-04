package main

import (
	"github.com/Hertucktor/archive-importer/config"
	"github.com/Hertucktor/archive-importer/importer"
	"github.com/Hertucktor/archive-importer/mongodb"
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
	client, err := mongodb.CreateClient(conf.DBUser, conf.DBPass, conf.DBPort, conf.DBName, logger)
	if err != nil {
		logger.Fatal(err)
	}

	if err = importer.ImportCardsIntoDatabase(client, conf, page, delimiter, logger); err != nil {
		logger.Fatal(err)
	}

}
