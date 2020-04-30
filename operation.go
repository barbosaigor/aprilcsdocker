package dockercs

import (
	"errors"
)

// ErrEmptyContainer implies empty container error
var ErrEmptyContainer = errors.New("Empty container")

// Stop container
func Stop(container string) error {
	if container == "" {
		return ErrEmptyContainer
	}
	id, err := getContainerID(container)
	if err != nil {
		return err
	}
	return stopContainer(id)
}

// StopAll stops all containers, but if a container
// throws an error then stop operation
func StopAll(container []string) error {
	if container == nil || len(container) == 0 {
		return ErrEmptyContainer
	}
	for _, c := range container {
		if err := Stop(c); err != nil {
			return err
		}
	}
	return nil
}
