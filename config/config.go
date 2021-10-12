package config

import (
	"github.com/spf13/viper"
)

var Configurations Config

type Config struct {
	Aws
	Redis
}

type Aws struct {
	Region   string
	Endpoint string
	Profile  string
	Dynamodb
}

type Dynamodb struct {
	Table string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	Db       int
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Configurations); err != nil {
		panic(err)
	}
}
