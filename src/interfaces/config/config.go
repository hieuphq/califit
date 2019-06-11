package config

import (
	"github.com/spf13/viper"
)

// Config contain configuration of db for migrator
// config var < env < command flag
type Config struct {
	AppName   string
	BaseURL   string
	JwtSecret string
	Port      string

	SQLHost     string
	SQLPort     string
	SQLUsername string
	SQLPassword string
	SQLDBName   string
}

// Reader get config from reader
type Reader interface {
	Read() (Config, error)
}

// Loader load config from reader into Viper
type Loader interface {
	Load(viper.Viper) (*viper.Viper, error)
}

// GetBy get config by reader(file or env)
func GetBy(r Reader) (Config, error) {
	return r.Read()
}

// GenerateConfigFromViper generate config from viper data
func GenerateConfigFromViper(v viper.Viper) Config {
	return Config{
		Port:      v.GetString("PORT"),
		AppName:   v.GetString("APP_NAME"),
		BaseURL:   v.GetString("BASE_URL"),
		JwtSecret: v.GetString("JWT_SECRET"),

		SQLHost:     v.GetString("SQL_HOST"),
		SQLPort:     v.GetString("SQL_PORT"),
		SQLUsername: v.GetString("SQL_USERNAME"),
		SQLPassword: v.GetString("SQL_PASSWORD"),
		SQLDBName:   v.GetString("SQL_DBNAME"),
	}
}

// DefaultConfigLoaders is default loader list
func DefaultConfigLoaders() []Loader {
	loaders := []Loader{}

	fileLoader, err := NewFileLoader("env", ".")

	if err == nil {
		loaders = append(loaders, fileLoader)
	}

	return loaders
}

// LoadConfig load config from loader list
func LoadConfig(loaders []Loader) Config {
	v := viper.New()
	v.SetDefault("PORT", "8080")
	v.SetDefault("ENV", "development")

	for idx := range loaders {
		newV, err := loaders[idx].Load(*v)

		if err == nil {
			v = newV
		}
	}
	return GenerateConfigFromViper(*v)
}
