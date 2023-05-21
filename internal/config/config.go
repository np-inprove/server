package config

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/lestrrat-go/jwx/v2/jwa"
)

var (
	k             = koanf.New(".")
	letters       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	defaultConfig = map[string]interface{}{
		"seed.rootemail":                "root@np-inprove.com",
		"seed.rootpassword":             randPassword(),
		"seed.rootinstitutionname":      "iNProve",
		"seed.rootinstitutionshortname": "inprove",
	}
)

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

// SeedRootEmail is the email to be used for the seeded root user.
func (Config) SeedRootEmail() string {
	return k.String("seed.rootemail")
}

// SeedRootPassword is the password to be used for the seeded root user.
func (Config) SeedRootPassword() string {
	return k.String("seed.rootpassword")
}

// SeedRootInstitutionName is the name of the institution
// that the seeded root user is in. This value will
// only be used if an institution with the short name
// specified in SeedRootInstitutionShortName does not exist
func (Config) SeedRootInstitutionName() string {
	return k.String("seed.rootinstitutionname")
}

// SeedRootInstitutionShortName is the short name of the institution
// that the seeded root user is in.
func (Config) SeedRootInstitutionShortName() string {
	return k.String("seed.rootinstitutionshortname")
}

// SeedForceCreate forces the root user to be recreated
// even if it already exists.
func (Config) SeedForceCreate() bool {
	return k.Bool("seed.forcecreate")
}

func New() (*Config, error) {
	_ = k.Load(confmap.Provider(defaultConfig, "."), nil)
	_ = k.Load(file.Provider("./config.yaml"), yaml.Parser())

	err := k.Load(env.Provider(envPrefix, ".", func(s string) string {
		return strings.Replace(
			strings.ToLower(strings.TrimPrefix(s, envPrefix)), "_", ".", -1,
		)
	}), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	c := &Config{}

	var a jwa.SignatureAlgorithm
	if err := a.Accept(c.AppJWTAlgorithm()); err != nil {
		return nil, fmt.Errorf("failed to parse jwt algorithm: %w", err)
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
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	c := &Config{}

	if !strings.Contains(c.DatabaseDataSourceName(), "test") {
		return nil, fmt.Errorf("database source name used for test cases should contain 'test' substring: %s", c.DatabaseDataSourceName())
	}

	return &Config{}, nil
}

func randPassword() string {
	b := make([]rune, 12)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
