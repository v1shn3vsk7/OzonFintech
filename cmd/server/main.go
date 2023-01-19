package main

import (
	"OzonTestTask/internal/app/server"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	cfg := server.NewConfig()
	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(cfg)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}


}

