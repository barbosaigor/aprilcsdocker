package main

import (
	"os"
	"strconv"

	"github.com/barbosaigor/april/destroyer"
	"github.com/barbosaigor/aprilcsdocker"
)

type dockerDestryr struct{}

func (d dockerDestryr) Destroy(nodes []string) error {
	return aprilcsdocker.StopAll(nodes)
}

func main() {
	port := 7071
	if len(os.Args) > 1 {
		p, err := strconv.Atoi(os.Args[1])
		if err != nil {
			return
		}
		port = p
	}
	serv := destroyer.New(port, dockerDestryr{})
	serv.Serve()
}
