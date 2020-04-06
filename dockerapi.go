package dockercs

import (
	"context"
	"errors"
	"fmt"

	"github.com/barbosaigor/april/destroyer"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func stopContainer(containerId string) error {
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

type Container struct {
	Id string
	Name string
	Status destroyer.Status
}

// listContainers list running containers, but if all is true
// stopped containers is listed too.
func listContainers(all bool) ([]Container, error) {
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

	ctns := make([]Container, len(containers))
	for i, container := range containers {
		ctns[i].Id = container.ID
		if len(container.Names) > 0 {
			ctns[i].Name = container.Names[0][1:]
		}
		// TODO: Status
	}
	return ctns, nil
}

// getContainersId get running containers names mapped to an id, but if all is true
// stopped containers is listed too.
func getContainersId(all bool) (map[string]string, error) {
	containers, err := listContainers(all)
	if err != nil {
		return nil, err
	}
	ctns := make(map[string]string, len(containers))
	for _, container := range containers {
		if container.Name != "" {
			ctns[container.Name] = container.Id
		} else {
			ctns[container.Id] = container.Id
		}
	}
	return ctns, nil
}

func getContainerId(container string) (string, error) {
	ctns, err := getContainersId(true)
	if err != nil {
		return "", err
	}

	cId, ok := ctns[container]
	if !ok {
		return "", errors.New("Container not found")
	}
	return cId, nil
}
