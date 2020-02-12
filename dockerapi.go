package aprilcsdocker

import (
	"context"
	"errors"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func stop(containerId string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("Unable to create docker client")
		return err
	}
	err = cli.ContainerStop(context.Background(), containerId, nil)
	if err != nil {
		fmt.Println("Unable to stop docker container")
		return err
	}
	return nil
}

// listContainers list running containers, but if all is true
// stopped containers is listed too.
func listContainers(all bool) (map[string]string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("Unable to create docker client")
		return nil, err
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: all})
	if err != nil {
		fmt.Println("Unable to list docker containers", err.Error())
		return nil, err
	}

	ctns := make(map[string]string, len(containers))
	for _, container := range containers {
		if len(container.Names) > 0 {
			ctns[container.Names[0][1:]] = container.ID
		} else {
			ctns[container.ID] = container.ID
		}
	}
	return ctns, nil
}

func getContainerId(container string) (string, error) {
	ctns, err := listContainers(true)
	if err != nil {
		return "", err
	}

	cId, ok := ctns[container]
	if !ok {
		return "", errors.New("Container not found")
	}
	return cId, nil
}
