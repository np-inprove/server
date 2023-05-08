package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	HTTP struct {
		Host string
		Port int
	}
	Database struct {
		Host     string
		Name     string
		Username string
		Password string
	}
	App struct {
		Production bool
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

	if !strings.Contains(config.Database.Name, "test") {
		return nil, fmt.Errorf("database name used for testing did not contain 'test' substring: %s", config.Database.Name)
	}

	return config, nil
}
