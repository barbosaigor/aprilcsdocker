package aprilcsdocker

import (
	"fmt"
	"testing"
)

func TestListContainer(t *testing.T) {
	ctns, err := listContainer()
	if err != nil {
		t.Errorf("Error to list containers.\n%v", err.Error())
	}
	fmt.Println("Containers: ", ctns)
}

func TestGetContainerId(t *testing.T) {
	_, err := listContainer()
	if err != nil {
		t.Errorf("Error to list containers.\n%v", err.Error())
	}
	// Test is created manually
	cId, err := getContainerId("test")
	if err != nil {
		t.Errorf("Error to get container id.\n%v", err.Error())
	}
	fmt.Println("Container ID: ", cId)
}
