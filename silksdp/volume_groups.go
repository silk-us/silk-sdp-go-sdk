package silksdp

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// CreateVolumeGroup creates a new Volume Group on the Silk server.
//
// `enableDeDuplication` corresponds to "Provisioning Type" in the UI. When set to true, the Provisioning Type will be "thin Pprovisioning with dedupe"
func (c *Credentials) CreateVolumeGroup(name string, quotaInGb int, enableDeDuplication bool, description string, capacityPolicy string, timeout ...int) (*CreateOrUpdateVolumeGroupResponse, error) {

	httpTimeout := httpTimeout(timeout)

	config := map[string]interface{}{}
	config["name"] = name
	config["quota"] = quotaInGb * 1024 * 1024
	config["is_dedupe"] = enableDeDuplication
	config["description"] = description
	config["capacityPolicy"] = capacityPolicy

	apiRequest, err := c.Post("/volume_groups", config, httpTimeout)
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
// Valid config keys are: name, quota, capacity_policy, and description.
func (c *Credentials) UpdateVolumeGroup(name string, config map[string]interface{}, timeout ...int) (*CreateOrUpdateVolumeGroupResponse, error) {
	httpTimeout := httpTimeout(timeout)

	// Validate that the user provided keys are valid for this API
	validUpdateKeys := []string{"name", "quota", "capacity_policy", "description"}
	var invalidUserProvidedKeys []string
	for key := range config {

		if c.stringInSlice(validUpdateKeys, key) == false {
			invalidUserProvidedKeys = append(invalidUserProvidedKeys, key)
		}
	}

	// Return an error message if any invalid keys are found
	if len(invalidUserProvidedKeys) != 0 {
		return nil, fmt.Errorf("The provided 'config' parameter contains invalid keys. 'name', 'quota', 'capacity_policy', and 'description' are the only valid choices")
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

	objectsOnServer, err := c.GetVolumeGroups(httpTimeout)
	if err != nil {
		return 0, err
	}

	// Set objectID to a value (-1) that can not be returned by the server
	objectID := -1
	for _, object := range objectsOnServer.Hits {
		if object.Name == name {
			objectID = object.ID
		}

	}

	// If the objectID has not been updated (i.e not found on the server) return an error message
	if objectID == -1 {
		return 0, fmt.Errorf("The server does not contain a Volume Group named '%s'", name)
	}

	return objectID, nil

}
