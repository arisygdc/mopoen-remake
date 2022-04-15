package config

import (
	"github.com/spf13/viper"
)

type Environment struct {
	DBDriver      string `mapstructure:"DATABASE_DRIVER"`
	DBSource      string `mapstructure:"DATABASE_SOURCE"`
	ServerEnv     string `mapstructure:"SERVER_ENVIRONMENT"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func New(path string) (env Environment, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&env)
	return
}
