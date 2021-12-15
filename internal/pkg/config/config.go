package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DSN  string `yaml:"dsn"`
	GRPC GRPC   `yaml:"grpc"`
}

type GRPC struct {
	Port string `yaml:"port"`
}

func Load(path string) (*Config, error) {
	conf := Config{}
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return &conf, nil

}
