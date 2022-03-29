package silksdp

import (
	"testing"
)

func Test_CreateVolume(t *testing.T) {
	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	volumeName := "TestVolumeName"
	size := 10
	volumeGroupName := "TestVolumeGroupName"
	vmware := false
	description := "Created through the Go SDK"
	readOnly := false

	createVolume, err := silk.CreateVolume(volumeName, size, volumeGroupName, vmware, description, readOnly)
	if err != nil {
		t.Errorf("Failed to create volume: %v", err)
	}
	t.Logf("Volume created: %v", createVolume)
	// fmt.Println(createVolume)
}

func Test_GetVolumes(t *testing.T) {
	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	getVolumes, err := silk.GetVolumes()
	if err != nil {
		t.Errorf("Failed to fetch all volumes: %v", err)
	}
	t.Logf("Volume list: %v", getVolumes)
}

func Test_UpdateVolume(t *testing.T) {
	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	volumeName := "TestVolumeName"

	config := map[string]interface{}{}
	config["description"] = "New description"
	// config["name"] = "NewTestVolumeName"

	updateVolume, err := silk.UpdateVolume(volumeName, config)
	if err != nil {
		t.Errorf("Failed to update volume: %v", err)
	}
	t.Logf("Volume updated: %v", updateVolume)
}

func Test_DeleteVolume(t *testing.T) {
	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	volumeName := "TestVolumeName"

	deleteVolume, err := silk.DeleteVolume(volumeName)
	if err != nil {
		t.Errorf("Failed to delete volume: %v", err)
	}
	t.Logf("Volume deleted: %v", deleteVolume)
}

func Test_LifecycleVolume(t *testing.T) {
	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	Test_CreateVolume(t)
	Test_GetVolumes(t)
	Test_UpdateVolume(t)
	Test_GetVolumes(t)
	Test_DeleteVolume(t)
}
