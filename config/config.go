package config

import "github.com/spf13/viper"

type Config struct {
	HttpServer HttpServer
	Postgres   Postgres
	Email      Email
}

type HttpServer struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
}

type Postgres struct {
	Host            string
	Port            string
	Login           string
	Password        string
	ConnMaxLifeTime int
	MaxOpenConn     int
	MaxIdleConn     int
}

type Email struct {
	SenderPassword string
}

func ReadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
