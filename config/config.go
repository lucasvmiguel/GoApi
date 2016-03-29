package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	path string
	Api  struct {
		Port int
	}
}

func New(path string) *Config {
	return &Config{
		path: path,
	}
}

func (c *Config) Read(env string) {
	path, err := os.Getwd()
	if err != nil {
		panic("Can't get path config file")
	}

	viper.SetConfigName(env)
	viper.AddConfigPath(path + "/" + c.path)

	if viper.ReadInConfig() != nil {
		panic("Can't read config file")
	}

	populateFields(c)
}

func populateFields(c *Config) {
	c.Api.Port = viper.GetInt("api.port")
}
