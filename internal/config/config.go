package config

import (
	"github.com/spf13/viper"
)

// Config is used to store data from config.yaml file in convenient struct
type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}
	DB struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
		Name string `mapstructure:"name"`
	}
	JWT struct {
		SecretKey string `mapstructure:"secret_key"`
	}
}

var AppConfig Config

// SetUp is used to set up global var for storing config data
func SetUp() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config/")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}
	return nil
}
