package main

import (
	"OzonTestTask/internal/app/server"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "/Users/vladimirvasilev/Desktop/Projects/OzonTestTask/configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	cfg := server.NewConfig()
	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		log.Fatal(err)
	}

	arg := os.Args[1:]

	if err := server.Start(cfg, arg[0]); err != nil {
		log.Fatal(err)
	}

}

