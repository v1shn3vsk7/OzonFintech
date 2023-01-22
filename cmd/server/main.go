package main

import (
	"OzonTestTask/internal/app/server"
	"log"
)

func main() {
	cfg := server.NewConfig()

	if err := server.Start(cfg); err != nil {
		log.Fatal(err)
	}

}
