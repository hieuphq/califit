package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// FileReader read config from file
type FileReader struct {
	filename string
	dirname  string
}

// NewFileReader create new file reader with filename and dirname
func NewFileReader(filename, dirname string) Reader {
	return &FileReader{filename, dirname}
}

// NewFileLoader create new file loader with filename and dirname
func NewFileLoader(filename, dirname string) (Loader, error) {
	filePath := fmt.Sprintf("%s/%s.yaml", dirname, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, err
	}
	return &FileReader{filename, dirname}, nil
}

// Load from yaml file
func (r *FileReader) Load(v viper.Viper) (*viper.Viper, error) {
	v.SetConfigType("yaml")
	v.SetConfigName(r.filename)
	v.AddConfigPath(r.dirname)

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *FileReader) Read() (Config, error) {
	v := viper.New()
	{
		v.SetConfigType("yaml")
		v.SetConfigName(r.filename)
		v.AddConfigPath(r.dirname)
	}

	err := v.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	return Config{
		JwtSecret:   v.GetString("JWT_SECRET"),
		SQLHost:     v.GetString("SQL_HOST"),
		SQLPort:     v.GetString("SQL_PORT"),
		SQLUsername: v.GetString("SQL_USERNAME"),
		SQLPassword: v.GetString("SQL_PASSWORD"),
		SQLDBName:   v.GetString("SQL_DBNAME"),
		Port:        v.GetString("PORT"),
		AppName:     v.GetString("APP_NAME"),
		BaseURL:     v.GetString("BASE_URL"),
	}, nil
}
