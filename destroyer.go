package dockercs

import "github.com/barbosaigor/april/destroyer"

// Implement chaos server interface from april/destroyer
type ChaosServerDocker struct{}

func (d ChaosServerDocker) Shutdown(svc string) error {
	if id, err := getContainerId(svc); err != nil {
		return err
	} else {
		return stopContainer(id)
	}
}

func (d ChaosServerDocker) ListInstances(status destroyer.Status) ([]destroyer.Instance, error) {
	containers, err := listContainers(status == destroyer.Any)
	if err != nil {
		return nil, err
	}
	svcs := make([]destroyer.Instance, len(containers))
	for i, c := range containers {
		svcs[i] = destroyer.Instance{c.Name, c.Status}
	}
	return svcs, nil
}