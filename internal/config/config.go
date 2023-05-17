package config

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"strings"

	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

const envPrefix = "INPROVE_"

type Config struct{}

func (Config) HTTPAddr() string {
	return k.String("http.addr")
}

func (Config) DatabaseDriverName() string {
	return k.String("database.drivername")
}

func (Config) DatabaseDataSourceName() string {
	return k.String("database.datasourcename")
}

func (Config) DatabaseAutoMigration() bool {
	return k.Bool("database.automigration")
}

func (Config) AppProduction() bool {
	return k.Bool("app.production")
}

func (Config) AppJWTAlgorithm() string {
	return k.String("app.jwtalgorithm")
}

func (Config) AppJWTSignKey() string {
	return k.String("app.jwtsignkey")
}

func (Config) AppJWTVerifyKey() string {
	return k.String("app.jwtverifykey")
}

func New() (*Config, error) {
	_ = k.Load(file.Provider("./config.yaml"), yaml.Parser())

	err := k.Load(env.Provider(envPrefix, ".", func(s string) string {
		return strings.Replace(
			strings.ToLower(strings.TrimPrefix(s, envPrefix)), "_", ".", -1,
		)
	}), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %v", err)
	}

	return &Config{}, nil
}

//func NewTest() (*Config, error) {
//	viper.SetConfigName("config.testing")
//	viper.SetConfigType("yaml")
//	viper.AddConfigPath("config")
//	viper.AddConfigPath("./")
//	viper.AddConfigPath("../config")
//	viper.AddConfigPath("../../config")
//
//	err := viper.ReadInConfig()
//	if err != nil {
//		return nil, err
//	}
//
//	var config *Config
//	err = viper.Unmarshal(&config)
//	if err != nil {
//		return nil, err
//	}
//
//	if !strings.Contains(config.Database.DataSourceName, "test") {
//		return nil, fmt.Errorf("database source name used for test cases should contain 'test' substring: %s", config.Database.DataSourceName)
//	}
//
//	return config, nil
//}
