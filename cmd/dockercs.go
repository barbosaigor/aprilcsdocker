package main

import (
	"fmt"

	"github.com/barbosaigor/dockercs"
	"github.com/barbosaigor/april/destroyer/cli"
)

func main() {
	// Change cli long description
	cli.RootCmd.Long = "Dockercs (Chaos server) shutdowns Docker containers via an API."
	cli.Cs = dockercs.ChaosServerDocker{}
	if err := cli.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
