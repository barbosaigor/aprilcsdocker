package main

import (
	"log"

	"github.com/barbosaigor/april/destroyer/cli"
	"github.com/barbosaigor/dockercs"
)

func main() {
	// Change cli long description
	cli.RootCmd.Long = "Dockercs (Chaos server) shutdowns Docker containers via an API."
	cli.Cs = dockercs.ChaosServerDocker{}
	if err := cli.Execute(); err != nil {
		log.Printf("Error: %v\n", err)
	}
}
