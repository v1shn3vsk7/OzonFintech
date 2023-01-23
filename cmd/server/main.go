package main

import (
	"OzonTestTask/internal/app/server"
	"OzonTestTask/internal/app/server/gRPC"
	"log"
)

func main() {
	cfg := server.NewConfig()

	if err := server.Start(cfg); err != nil {
		log.Fatal(err)
	}

	if err := gRPC.Start(); err != nil {
		log.Fatal(err)
	}

}
