package config

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	HTTP struct {
		ListenAddr string
	}
	PostgreSQL struct {
		Username string
		Password string
		Host     string
		Database string
	}
	App struct {
		Environment string
	}
}

func New() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config *Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func NewTest() (*Config, error) {
	viper.SetConfigName("config.testing")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config *Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	if !strings.Contains(config.PostgreSQL.Database, "test") {
		return nil, errors.New("postgresql database for tests must contain 'test' in name")
	}

	return config, nil
}
