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
		DriverName     string
		DataSourceName string
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

	if !strings.Contains(config.Database.DataSourceName, "test") {
		return nil, fmt.Errorf("database source name used for test cases should contain 'test' substring: %s", config.Database.DataSourceName)
	}

	return config, nil
}
