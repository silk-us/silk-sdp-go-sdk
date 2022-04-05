package silksdp

import (
	"fmt"
	"testing"
)

func _BuildHostsAndHostGroups(t *testing.T, id int) {
	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	var hosts []string
	for i := 1; i < 5; i++ {
		hostname := fmt.Sprintf("example%d_test_host%02d", id, i)
		_, err := silk.CreateHost(hostname, "Linux")
		if err != nil {
			t.Errorf("Failed to create host: %v\n%v", hostname, err)
		} else {
			t.Logf("Host created: %v", hostname)
			hosts = append(hosts, hostname)
		}
		iqn := fmt.Sprintf("iqn.2022-03.com.example:c5f336d488%02d", i)
		_, err = silk.CreateHostIQN(hostname, iqn)
		if err != nil {
			t.Errorf("Failed to set iqn host: %v\n%v", hostname, err)
		} else {
			t.Logf("Host IQN set: %v", hostname)
		}
	}

	hostGroupName := fmt.Sprintf("example%d_test_hostgroup", id)
	_, err = silk.CreateHostGroup(hostGroupName, "this is a test group for test hosts", true)
	if err != nil {
		t.Errorf("Failed to create host group: %v", err)
	}
	t.Logf("Host group created: %v", hostGroupName)

	for _, hostname := range hosts {
		_, err = silk.CreateHostHostGroupMapping(hostname, hostGroupName)
		if err != nil {
			t.Errorf("Failed to add %v host group %v:\n%v", hostname, hostGroupName, err)
		} else {
			t.Logf("Add host to hostgroup: %v -> %v", hostname, hostGroupName)
		}
	}
}

func _DeleteHosts(t *testing.T, id int) {
	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	for i := 1; i < 5; i++ {
		hostname := fmt.Sprintf("example%d_test_host%02d", id, i)
		_, err := silk.DeleteHost(hostname)
		if err != nil {
			t.Errorf("Failed to delete host: %v\n%v", hostname, err)
		} else {
			t.Logf("Delete Host: %v", hostname)
		}
	}
}

func _DeleteHostGroups(t *testing.T, id int) {
	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}
	hostGroupName := fmt.Sprintf("example%d_test_hostgroup", id)

	_, err = silk.DeleteHostGroup(hostGroupName)
	if err != nil {
		t.Errorf("Failed to remove hostgroup: %v", hostGroupName)
		t.Errorf("%v", err)
	} else {
		t.Logf("Delete hostgroup %v", hostGroupName)
	}
}

func _CreateVolumesAndVolumeGroup(t *testing.T, id int) {
	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	// Create Volume Group
	volumeGroupName := fmt.Sprintf("example%d_test_volumegroup", id)
	quotaInGb := 10
	enableDedup := true
	description := "Created through the Silk Go SDK"
	capacityPolicy := "default_vg_capacity_policy"
	createNewVolumeGroup, err := silk.CreateVolumeGroup(volumeGroupName, quotaInGb, enableDedup, description, capacityPolicy)
	if err != nil {
		t.Errorf("Failed to create volume group: %v", err)
	}
	t.Logf("Volume group created: %v", createNewVolumeGroup)

	//Create Volumes
	for i := 1; i < 4; i++ {
		volumeName := fmt.Sprintf("example%d_test_volume%02d", id, i)
		size := 2
		vmware := false
		description := "Created through the Go SDK"
		readOnly := false

		_, err := silk.CreateVolume(volumeName, size, volumeGroupName, vmware, description, readOnly)
		if err != nil {
			t.Errorf("Failed to create volume: %v", err)
		}
		t.Logf("Volume created: %v (size:%dGB)", volumeName, size)
	}
}

func _MapVolumeGroupToHostGroup(t *testing.T, id int) {
	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	volumeGroupName := fmt.Sprintf("example%d_test_volumegroup", id)
	hostGroupName := fmt.Sprintf("example%d_test_hostgroup", id)

	hostGroupVolumeMapping, err := silk.CreateHostGroupVolumeGroupMapping(hostGroupName, volumeGroupName)
	if err != nil {
		t.Errorf("Failed to map %v to %v\n%v", volumeGroupName, hostGroupName, err)
	} else {
		t.Logf("Map %v to %v", volumeGroupName, hostGroupName)
		t.Logf("%v", hostGroupVolumeMapping)
	}
}

func _DeleteHostGroupMappings(t *testing.T, id int) {
	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}
	hostGroupName := fmt.Sprintf("example%d_test_hostgroup", id)

	_, err = silk.DeleteHostGroupMappings(hostGroupName)
	if err != nil {
		t.Errorf("Failed to remove host group mappings: %v", hostGroupName)
		t.Errorf("%v", err)
	} else {
		t.Logf("Delete host group mappings %v", hostGroupName)
	}
}

func _DeleteVolumeGroups(t *testing.T, id int) {
	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}
	volumeGroupName := fmt.Sprintf("example%d_test_volumegroup", id)

	_, err = silk.DeleteVolumeGroup(volumeGroupName)
	if err != nil {
		t.Errorf("Failed to remove volume group: %v", volumeGroupName)
		t.Errorf("%v", err)
	} else {
		t.Logf("Delete volume group %v", volumeGroupName)
	}
}

func _DeleteVolumes(t *testing.T, id int) {
	silk, err := ConnectEnv()
	if err != nil {
		t.Errorf("Failed to connect: %v", err)
	}

	for i := 1; i < 4; i++ {
		volumeName := fmt.Sprintf("example%d_test_volume%02d", id, i)
		_, err := silk.DeleteVolume(volumeName)
		if err != nil {
			t.Errorf("Failed to delete volume: %v", volumeName)
			t.Errorf("%v", err)
		} else {
			t.Logf("Delete volume: %v", volumeName)
		}
	}
}

func _DeleteHostsAndHostGroups(t *testing.T, id int) {
	_DeleteHosts(t, id)
	_DeleteHostGroups(t, id)
}

// Add and remove hosts
// Remove hostgroup before removing its hosts
func Test_Example1(t *testing.T) {
	_BuildHostsAndHostGroups(t, 1)
	t.Logf("BUILD COMPLETE")
	_DeleteHostGroups(t, 1)
	_DeleteHosts(t, 1)
	t.Logf("DESTROY COMPLETE")
}

// Add hostgroup with hosts and connect volume group with multiple volumes
// Remove everything in correct order (without unmapping)
func Test_Example2(t *testing.T) {
	_BuildHostsAndHostGroups(t, 2)
	_CreateVolumesAndVolumeGroup(t, 2)
	_MapVolumeGroupToHostGroup(t, 2)
	t.Logf("BUILD COMPLETE")
	_DeleteVolumes(t, 2)
	_DeleteVolumeGroups(t, 2)
	_DeleteHostGroups(t, 2)
	_DeleteHosts(t, 2)
	t.Logf("DESTROY COMPLETE")
}

// Add hostgroup with hosts and connect volume group with multiple volumes
// Remove hostgroup before hosts (unmap required)
func Test_Example3(t *testing.T) {
	_BuildHostsAndHostGroups(t, 3)
	_CreateVolumesAndVolumeGroup(t, 3)
	_MapVolumeGroupToHostGroup(t, 3)
	t.Logf("BUILD COMPLETE")
	_DeleteHostGroupMappings(t, 3)
	_DeleteHostGroups(t, 3)
	_DeleteHosts(t, 3)
	_DeleteVolumes(t, 3)
	_DeleteVolumeGroups(t, 3)
	t.Logf("DESTROY COMPLETE")
}

// Mount host to volume
func Test_Example4(t *testing.T) {
	_BuildHostsAndHostGroups(t, 4)

}
