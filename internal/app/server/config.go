package server

import "OzonTestTask/internal/app/data/inmemory"

type Config struct {
	BindAddr string
	LogLevel string
	DbUrl    string
	Data     *inmemory.Data
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		DbUrl:    "postgres://user:password@db:5432/ozontesttask?sslmode=disable",
		Data: &inmemory.Data{},
	}
}
