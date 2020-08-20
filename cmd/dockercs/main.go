package main

import (
	"log"

	"github.com/barbosaigor/april/chaosserver/cli"
	"github.com/barbosaigor/dockercs"
)

func main() {
	// Change cli long description
	cli.RootCmd.Long = "Dockercs (Chaos Server) shutdowns Docker containers."
	cli.Cs = dockercs.ChaosServerDocker{}
	if err := cli.Execute(); err != nil {
		log.Printf("Error: %v\n", err)
	}
}
