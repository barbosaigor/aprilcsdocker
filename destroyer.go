package dockercs

import "github.com/barbosaigor/april/destroyer"

// ChaosServerDocker implements chaos server interface from april/destroyer
type ChaosServerDocker struct{}

// Shutdown turns off a service
func (d ChaosServerDocker) Shutdown(svc string) error {
	id, err := getContainerID(svc)
	if err != nil {
		return err
	}
	return stopContainer(id)
}

// ListInstances lists all instances with corresponding status
func (d ChaosServerDocker) ListInstances(status destroyer.Status) ([]destroyer.Instance, error) {
	containers, err := listContainers(status == destroyer.Any)
	if err != nil {
		return nil, err
	}
	svcs := make([]destroyer.Instance, len(containers))
	for i, c := range containers {
		svcs[i] = destroyer.Instance{Name: c.Name, Sts: destroyer.Up}
	}
	return svcs, nil
}

// OnStart is a life cycle routine when stating the chaos server
func (d ChaosServerDocker) OnStart() {
}
