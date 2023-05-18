package config

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"strings"

	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

const envPrefix = "INPROVE_"

type Config struct{}

// HTTPAddr is the address that the server will listen on.
func (Config) HTTPAddr() string {
	return k.String("http.addr")
}

// DatabaseDriverName is the name of the database driver.
// This must be supported by ent and can only be `postgres` or `sqlite`.
func (Config) DatabaseDriverName() string {
	return k.String("database.drivername")
}

// DatabaseDataSourceName is the connection string for the
// chosen database driver.
func (Config) DatabaseDataSourceName() string {
	return k.String("database.datasourcename")
}

// DatabaseAutoMigration indicates whether ent should run auto migration
// on the database.
func (Config) DatabaseAutoMigration() bool {
	return k.Bool("database.automigration")
}

// AppProduction indicates whether the server should run in production mode
func (Config) AppProduction() bool {
	return k.Bool("app.production")
}

// AppJWTAlgorithm specifies the algorithm to use for signing the JWT used for auth
func (Config) AppJWTAlgorithm() jwa.SignatureAlgorithm {
	var a jwa.SignatureAlgorithm
	// This should not be error as the value must be checked when creating Config
	_ = a.Accept(k.String("app.jwtalgorithm"))
	return a
}

// AppJWK is the key used in AppJWTAlgorithm. The same algorithm must be used for both.
func (Config) AppJWK() string {
	return k.String("app.jwk")
}

func (Config) AppJWTCookieHost() string {
	return k.String("app.jwtcookiehost")
}

func (Config) AppJWTCookieName() string {
	return k.String("app.jwtcookiename")
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

	c := &Config{}

	var a jwa.SignatureAlgorithm
	if err := a.Accept(c.AppJWTAlgorithm()); err != nil {
		return nil, fmt.Errorf("failed to parse jwt algorithm: %v", err)
	}

	return c, nil
}

func NewTest() (*Config, error) {
	_ = k.Load(file.Provider("./config.testing.yaml"), yaml.Parser())

	err := k.Load(env.Provider(envPrefix, ".", func(s string) string {
		return strings.Replace(
			strings.ToLower(strings.TrimPrefix(s, envPrefix)), "_", ".", -1,
		)
	}), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %v", err)
	}

	c := &Config{}

	if !strings.Contains(c.DatabaseDataSourceName(), "test") {
		return nil, fmt.Errorf("database source name used for test cases should contain 'test' substring: %s", c.DatabaseDataSourceName())
	}

	return &Config{}, nil
}
