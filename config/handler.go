package config

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DBUser                string `yaml:"dbUser"`
	DBPass                string `yaml:"dbPass"`
	DBPort                string `yaml:"dbPort"`
	DBName                string `yaml:"dbName"`
	DBCollectionAllcards  string `yaml:"dbCollectionAllcards"`
	DBCollectionMycards   string `yaml:"dbCollectionMycards"`
	DBCollectionSetimages string `yaml:"dbCollectionSetimages"`
	DBCollectionSetNames  string `yaml:"dbCollectionSetNames"`
}

func GetConfig(configFile string, logger *zap.SugaredLogger) (Config, error) {
	var c Config

	buf, err := os.ReadFile(configFile)
	if err != nil {
		logger.Error(err)
		return c, err
	}

	if err = yaml.Unmarshal(buf, &c); err != nil {
		logger.Error(err)
		return c, err
	}

	return c, err
}
