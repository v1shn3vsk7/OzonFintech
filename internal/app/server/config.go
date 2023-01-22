package server

type Config struct {
	BindAddr string
	LogLevel string
	DbUrl    string
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		DbUrl:    "postgres://user:password@db:5432/ozontesttask?sslmode=disable",
	}
}
