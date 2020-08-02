package silksdp_test

import (
	"fmt"
	"log"

	"github.com/silk-us/silk-sdp-go-sdk/silksdp"
)

func ExampleCredentials_CreateVolumeGroup() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	volumeGroupName := "ExampleVolumeGroupName"
	quotaInGb := 10
	enableDedup := true
	description := "Created through the Silk Go SDK"
	capacityPolicy := "default_vg_capacity_policy"

	createNewVolumeGroup, err := silk.CreateVolumeGroup(volumeGroupName, quotaInGb, enableDedup, description, capacityPolicy)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(createNewVolumeGroup)

}
func ExampleCredentials_GetVolumeGroups() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	volumeGroups, err := silk.GetVolumeGroups()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(volumeGroups)

}

func ExampleCredentials_UpdateVolumeGroup() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	volumeGroupName := "ExampleVolumeGroupName"

	config := map[string]interface{}{}
	config["name"] = "UpdatedVolumeGroupName"
	config["quota"] = 20971520 // Value provided in kilobytes
	config["description"] = "Updated description"
	config["capacityPolicy"] = "new-vg-cap-policy"

	updateVolumeGroup, err := silk.UpdateVolumeGroup(volumeGroupName, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updateVolumeGroup)

}

func ExampleCredentials_DeleteVolumeGroup() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	volumeGroupName := "ExampleVolumeGroupName"

	deleteVolumeGroup, err := silk.DeleteVolumeGroup(volumeGroupName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deleteVolumeGroup)

}

func ExampleCredentials_CreateVolume() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	volumeName := "ExampleVolumeName"
	size := 10
	volumeGroupName := "ExampleVolumeGroupName"
	vmware := false
	description := "Created through the Go SDK"
	readOnly := false

	createVolume, err := silk.CreateVolume(volumeName, size, volumeGroupName, vmware, description, readOnly)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(createVolume)

}

func ExampleCredentials_GetVolumes() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	getVolume, err := silk.GetVolumes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getVolume)

}

func ExampleCredentials_UpdateVolume() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	volumeName := "ExampleVolumeName"

	config := map[string]interface{}{}
	config["name"] = "NewExampleVolumeName"

	updateVolume, err := silk.UpdateVolume(volumeName, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updateVolume)

}

func ExampleCredentials_DeleteVolume() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	volumeName := "ExampleVolumeName"

	deleteVolume, err := silk.DeleteVolume(volumeName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deleteVolume)

}

func ExampleCredentials_CreateHost() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	name := "ExampleHostName"
	hostType := "Linux"

	createHost, err := silk.CreateHost(name, hostType)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(createHost)

}

func ExampleCredentials_GetHosts() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	getHosts, err := silk.GetHosts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getHosts)

}

func ExampleCredentials_UpdateHost() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"

	config := map[string]interface{}{}
	config["name"] = "NewExampleHostName"
	config["type"] = "Windows"

	updateHost, err := silk.UpdateHost(hostName, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updateHost)

}

func ExampleCredentials_UpdateHost_addtohostgroup() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"
	hostGroupToAddTo := "ExampleHostGroupName"

	hostGroupID, err := silk.GetHostGroupID(hostGroupToAddTo)

	addToHostGroupConfig := map[string]string{}
	addToHostGroupConfig["ref"] = fmt.Sprintf("/host_groups/%d", hostGroupID)

	config := map[string]interface{}{}
	config["host_group"] = "addToHostGroupConfig"

	updateHost, err := silk.UpdateHost(hostName, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updateHost)

}

func ExampleCredentials_DeleteHost() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"

	deleteHost, err := silk.DeleteHost(hostName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deleteHost)

}

func ExampleCredentials_CreateHostGroup() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	name := "ExampleHostGroupName"
	description := "Created through the Go SDK"
	allowDifferentHostTypes := true

	createHostGroup, err := silk.CreateHostGroup(name, description, allowDifferentHostTypes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(createHostGroup)

}

func ExampleCredentials_GetHostGroups() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	getHostGroups, err := silk.GetHostGroups()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getHostGroups)

}

func ExampleCredentials_UpdateHostGroup() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostGroupName := "ExampleHostGroupName"

	config := map[string]interface{}{}
	config["description"] = "New Example Description"
	config["allow_different_host_types"] = false

	updateHostGroup, err := silk.UpdateHostGroup(hostGroupName, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updateHostGroup)

}

func ExampleCredentials_DeleteHostGroup() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostGroupName := "ExampleHostGroupName"

	deleteHostGroup, err := silk.DeleteHostGroup(hostGroupName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deleteHostGroup)

}

func ExampleCredentials_CreateHostVolumeMapping() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"
	volumeName := "ExampleVolumeName"

	createHostVolumeMapping, err := silk.CreateHostVolumeMapping(hostName, volumeName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(createHostVolumeMapping)

}

func ExampleCredentials_GetHostMappings() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	getHostMappings, err := silk.GetHostMappings()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getHostMappings)

}

func ExampleCredentials_DeleteHostMappings() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"

	deleteAllHostMappings, err := silk.DeleteHostMappings(hostName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deleteAllHostMappings)

}

func ExampleCredentials_DeleteHostVolumeMapping() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"
	volumeName := "ExampleVolumeName"

	deleteSingeVolumeHostMapping, err := silk.DeleteHostVolumeMapping(hostName, volumeName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deleteSingeVolumeHostMapping)

}

func ExampleCredentials_GetHostGroupID() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostGroupName := "ExampleHostGroupName"

	hostGroupID, err := silk.GetHostGroupID(hostGroupName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hostGroupID)

}

func ExampleCredentials_GetVolumeID() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	volumeName := "ExampleVolumeName"

	volumeID, err := silk.GetVolumeID(volumeName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(volumeID)

}

func ExampleCredentials_GetHostID() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"

	hostID, err := silk.GetHostID(hostName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hostID)

}

func ExampleCredentials_GetVolumeGroupID() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	volumeGroupName := "ExampleVolumeGroupName"

	volumeGroupID, err := silk.GetVolumeGroupID(volumeGroupName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(volumeGroupID)

}

func ExampleCredentials_CreateHostGroupVolumeMapping() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostGroupName := "ExampleHostGroupName"
	volumeName := "ExampleVolumeName"

	createHostGroupVolumeMapping, err := silk.CreateHostGroupVolumeMapping(hostGroupName, volumeName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(createHostGroupVolumeMapping)

}

func ExampleCredentials_GetHostGroupMappings() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	getHostGroupMappings, err := silk.GetHostGroupMappings()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getHostGroupMappings)

}

func ExampleCredentials_DeleteHostGroupMappings() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostGroupName := "ExampleHostGroupName"

	deleteHostGroupMappings, err := silk.DeleteHostGroupMappings(hostGroupName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deleteHostGroupMappings)

}

func ExampleCredentials_DeleteHostGroupVolumeMapping() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostGroupName := "ExampleHostGroupName"
	volumeName := "ExampleVolumeName"

	deleteHostGroupVolumeMapping, err := silk.DeleteHostGroupVolumeMapping(hostGroupName, volumeName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deleteHostGroupVolumeMapping)

}

func ExampleCredentials_CreateHostPWWN() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"
	pwwn := "20:16:33:79:55:99:ab:9f"

	createHostPWWN, err := silk.CreateHostPWWN(hostName, pwwn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(createHostPWWN)

}

func ExampleCredentials_GetHostPWWN() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"

	getHostPWWN, err := silk.GetHostPWWN(hostName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getHostPWWN)

}

func ExampleCredentials_DeleteHostPWWN() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"

	deleteHostPWWN, err := silk.DeleteHostPWWN(hostName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deleteHostPWWN)

}

func ExampleCredentials_GetHostIQN() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"

	getHostIQN, err := silk.GetHostIQN(hostName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getHostIQN)

}

func ExampleCredentials_DeleteHostIQN() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"

	deleteHostIQN, err := silk.DeleteHostIQN(hostName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deleteHostIQN)

}

func ExampleCredentials_CreateHostIQN() {

	// Use ConnectEnv to look up the Silk Server, Username, and Password
	// using environment variables
	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	hostName := "ExampleHostName"
	iqn := "iqn.2009-01.com.kaminario:storage.k2.2289`"

	createHostIQN, err := silk.CreateHostIQN(hostName, iqn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(createHostIQN)

}
