package dockercs

import "github.com/barbosaigor/april/chaosserver"

// ChaosServerDocker implements chaos server interface from april/chaosserver
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
func (d ChaosServerDocker) ListInstances(status chaosserver.Status) ([]chaosserver.Instance, error) {
	containers, err := listContainers(status == chaosserver.Any)
	if err != nil {
		return nil, err
	}
	svcs := make([]chaosserver.Instance, len(containers))
	for i, c := range containers {
		svcs[i] = chaosserver.Instance{Name: c.Name, Sts: chaosserver.Up}
	}
	return svcs, nil
}

// OnStart is a life cycle routine when stating the chaos server
func (d ChaosServerDocker) OnStart() {
}
