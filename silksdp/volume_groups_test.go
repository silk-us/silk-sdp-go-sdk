package silksdp

import (
	"testing"
)

func Test_CreateVolumeGroup(t *testing.T) {

	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	volumeGroupName := "TestVolumeGroupName"
	quotaInGb := 10
	enableDedup := true
	description := "Created through the Silk Go SDK"
	capacityPolicy := "default_vg_capacity_policy"

	createNewVolumeGroup, err := silk.CreateVolumeGroup(volumeGroupName, quotaInGb, enableDedup, description, capacityPolicy)
	if err != nil {
		t.Errorf("Failed to create volume group: %v", err)
	}
	t.Logf("Volume group created: %v", createNewVolumeGroup)
	// fmt.Println(createVolume)
}

func Test_GetVolumeGroups(t *testing.T) {

	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	volumeGroups, err := silk.GetVolumeGroups()
	if err != nil {
		t.Errorf("Failed to fetch volume groups: %v", err)
	}
	t.Logf("Volume groups: %v", volumeGroups)

}

func Test_UpdateVolumeGroup(t *testing.T) {

	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	volumeGroupName := "TestVolumeGroupName"

	config := map[string]interface{}{}
	// config["name"] = "UpdatedVolumeGroupName"
	config["quotaInGb"] = 11 // Value provided in kilobytes
	config["description"] = "Updated description"
	// config["capacityPolicy"] = "new-vg-cap-policy"

	updateVolumeGroup, err := silk.UpdateVolumeGroup(volumeGroupName, config)
	if err != nil {
		t.Errorf("Failed to update volume group: %v", err)
	}
	t.Logf("Volume group updated: %v", updateVolumeGroup)

}

func Test_DeleteVolumeGroup(t *testing.T) {

	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	volumeGroupName := "TestVolumeGroupName"

	deleteVolumeGroup, err := silk.DeleteVolumeGroup(volumeGroupName)
	if err != nil {
		t.Errorf("Failed to delete volume group: %v", err)
	}
	t.Logf("Volume group deleted: %v", deleteVolumeGroup)

}

func Test_LifecycleVolumeGroup(t *testing.T) {
	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	Test_CreateVolumeGroup(t)
	Test_GetVolumeGroups(t)
	Test_UpdateVolumeGroup(t)
	Test_GetVolumeGroups(t)
	Test_DeleteVolumeGroup(t)
}