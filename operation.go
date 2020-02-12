package dockercs

import (
	"errors"
)

var EmptyContainer = errors.New("Empty container")

// Stop container
func Stop(container string) error {
	if container == "" {
		return EmptyContainer
	}
	id, err := getContainerId(container)
	if err != nil {
		return err
	}
	return stop(id)
}

// StopAll stops all containers, but if a container
// throws an error then stop operation
func StopAll(container []string) error {
	if container == nil {
		return EmptyContainer
	}
	for _, c := range container {
		if err := Stop(c); err != nil {
			return err
		}
	}
	return nil
}
