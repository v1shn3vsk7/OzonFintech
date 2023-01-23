package gRPC

type Config struct {
	BindAddr string
	DbUrl    string
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8433",
		DbUrl:    "postgres://user:password@db:5432/ozontesttask?sslmode=disable",
	}
}
