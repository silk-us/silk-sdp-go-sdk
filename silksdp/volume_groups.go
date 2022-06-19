package silksdp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// CreateVolumeGroup creates a new Volume Group on the Silk server.
//
// `enableDeDuplication` corresponds to "Provisioning Type" in the UI. When set to true, the Provisioning Type will be "thin provisioning with dedupe"
func (c *Credentials) CreateVolumeGroup(name string, quotaInGb int, enableDeDuplication bool, description string, capacityPolicy string, timeout ...int) (*CreateOrUpdateVolumeGroupResponse, error) {

	httpTimeout := httpTimeout(timeout)

	config := map[string]interface{}{}
	config["name"] = name
	config["quota"] = quotaInGb * 1024 * 1024
	config["is_dedup"] = enableDeDuplication
	config["description"] = description
	config["capacityPolicy"] = capacityPolicy

	apiRequest, err := c.Post("/volume_groups", config, httpTimeout)
	if err != nil {
		return nil, err
	}

	var apiResponse CreateOrUpdateVolumeGroupResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// GetVolumeGroups returns information on all Volume Groups found on the Silk server.
func (c *Credentials) GetVolumeGroups(timeout ...int) (*GetVolumeGroupsResponse, error) {

	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.Get("/volume_groups", httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse GetVolumeGroupsResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// UpdateVolumeGroup updates the Volume Group with the provided config options.
//
// Valid config keys are: name, quota, capacityPolicy, and description.
func (c *Credentials) UpdateVolumeGroup(name string, config map[string]interface{}, timeout ...int) (*CreateOrUpdateVolumeGroupResponse, error) {
	httpTimeout := httpTimeout(timeout)

	if _, ok := config["quotaInGb"]; ok {
		config["quota"] = config["quotaInGb"].(int) * 1024 * 1024
		delete(config, "quotaInGb")
	}

	// Validate that the user provided keys are valid for this API
	validUpdateKeys := []string{"name", "quota", "capacityPolicy", "description"}
	var invalidUserProvidedKeys []string
	for key := range config {

		if c.stringInSlice(validUpdateKeys, key) == false {
			invalidUserProvidedKeys = append(invalidUserProvidedKeys, key)
		}
	}

	// Return an error message if any invalid keys are found
	if len(invalidUserProvidedKeys) != 0 {
		return nil, fmt.Errorf("The provided 'config' parameter contains invalid keys. 'name', 'quota', 'capacityPolicy', and 'description' are the only valid choices")
	}

	volumeGroupID, err := c.GetVolumeGroupID(name, httpTimeout)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Patch(fmt.Sprintf("/volume_groups/%d", volumeGroupID), config, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateOrUpdateVolumeGroupResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil

}

// DeleteVolumeGroup deletes a Volume Group from the Silk server.
func (c *Credentials) DeleteVolumeGroup(name string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	volumeGroupID, err := c.GetVolumeGroupID(name, httpTimeout)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Delete(fmt.Sprintf("/volume_groups/%d", volumeGroupID), httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse DeleteResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil

}

// GetVolumeGroupID provides the ID for the provided Volume Group name.
func (c *Credentials) GetVolumeGroupID(name string, timeout ...int) (int, error) {

	httpTimeout := httpTimeout(timeout)

	allVolumeGroups, err := c.GetVolumeGroups(httpTimeout)
	if err != nil {
		return 0, err
	}

	// Set volumeGroupID to a value (-1) that can not be returned by the server
	volumeGroupID := -1
	for _, volumeGroup := range allVolumeGroups.Hits {
		if volumeGroup.Name == name {
			volumeGroupID = volumeGroup.ID
		}

	}

	// If the volumeGroupID has not been updated (i.e not found on the server) return an error message
	if volumeGroupID == -1 {
		return 0, fmt.Errorf("The server does not contain a Volume Group named '%s'", name)
	}

	return volumeGroupID, nil

}

// GetCapacityPolicyName returns the name of the Capacity Police based on the provided Capacity Policy id.
func (c *Credentials) GetCapacityPolicyName(id int, timeout ...int) (string, error) {

	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.Get("/vg_capacity_policies", httpTimeout)
	if err != nil {
		return "", err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse GetCapacityPolicyResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return "", mapErr
	}

	capacityPolicyName := ""
	for _, capacityPolicy := range apiResponse.Hits {
		if capacityPolicy.ID == id {
			capacityPolicyName = capacityPolicy.Name
		}
	}

	// If the capacityPolicyName has not been updated (i.e not found on the server) return an error message
	if capacityPolicyName == "" {
		return "", fmt.Errorf("The server does not contain a Capacity Policy with the ID of '%d'", id)
	}

	return capacityPolicyName, nil
}

// GetVolumeGroupHostMappings returns all Hosts that are mapped to the provided Volume Group.
func (c *Credentials) GetVolumeGroupHostMappings(volumeGroupName string, timeout ...int) ([]string, error) {

	httpTimeout := httpTimeout(timeout)

	volumeGroupID, err := c.GetVolumeGroupID(volumeGroupName)
	if err != nil {
		return nil, err
	}

	hostMappingsOnServer, err := c.GetHostMappings(httpTimeout)
	if err != nil {
		return nil, err
	}

	// Filter out the user provided volume and host from the hostMappingsOnServer
	// results
	hostName := []string{}

	for _, mapping := range hostMappingsOnServer {
		if mapping.Volume.Ref == fmt.Sprintf("/volume_groups/%d", volumeGroupID) {
			var hostID int
			if strings.Contains(mapping.Host.Ref, "/hosts") == true {

				hostRefID := strings.Replace(mapping.Host.Ref, "/hosts/", "", 1)
				hostID, err = strconv.Atoi(hostRefID)
				if err != nil {
					return nil, err
				}

				name, err := c.GetHostName(hostID)
				if err != nil {
					return nil, err
				}

				hostName = append(hostName, name)
			}

		}

	}

	// // If the mappingID has not been updated (i.e not found on the server) return an error message
	// if len(hostName) == 0 {
	// 	return nil, fmt.Errorf("No Host Mappings found on the Volume Group '%s'", volumeGroupName)
	// }

	return hostName, nil
}

// GetVolumeGroupVolumes provides the name of every Volume in a Volume Group.
func (c *Credentials) GetVolumeGroupVolumes(name string, timeout ...int) ([]string, error) {

	httpTimeout := httpTimeout(timeout)

	volumeGroupID, err := c.GetVolumeGroupID(name)
	if err != nil {
		return nil, err
	}

	allVolumes, err := c.GetVolumes(httpTimeout)
	if err != nil {
		return nil, err
	}

	volumes := []string{}
	for _, volume := range allVolumes.Hits {
		if string(volume.VolumeGroup.Ref) == fmt.Sprintf("/volume_groups/%d", volumeGroupID) {
			volumes = append(volumes, volume.Name)
		}
	}

	// if len(volumes) == 0 {
	// 	return nil, fmt.Errorf("The Volume Group '%s' does not contain any Volumes", name)
	// }

	return volumes, nil

}
