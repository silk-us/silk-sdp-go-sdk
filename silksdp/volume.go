package silksdp

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// CreateVolume creates a new Volume on the Silk server.
//
// The `volumeGroupName` corresponds to which Volume Group you wish to add the volume to.
// `vmware` corresponds to the "VMware Support" checkbox in the UI.
// `readOnly` corresponds to the "Exposure Type" radio button in the UI. When set to false, which is the default UI option, the volume will be set
// set to "Read/Only"
func (c *Credentials) CreateVolume(name string, sizeInGb int, volumeGroupName string, vmware bool, description string, readOnly bool, timeout ...int) (*CreateOrUpdateVolumeResponse, error) {

	httpTimeout := httpTimeout(timeout)

	volumeGroupID, err := c.GetVolumeGroupID(volumeGroupName)
	if err != nil {
		return nil, err
	}

	volumeGroupConfig := map[string]interface{}{}
	volumeGroupConfig["ref"] = fmt.Sprintf("/volume_groups/%d", volumeGroupID)

	config := map[string]interface{}{}
	config["name"] = name
	config["size"] = sizeInGb * 1024 * 1024
	config["volume_group"] = volumeGroupConfig
	config["vmware_support"] = vmware
	config["description"] = description
	config["read_only"] = readOnly

	apiRequest, err := c.Post("/volumes", config, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateOrUpdateVolumeResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// GetVolumes returns information on all Volumes found on the Silk server.
func (c *Credentials) GetVolumes(timeout ...int) (*GetVolumesResponse, error) {

	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.Get("/volumes", httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var createVolume GetVolumesResponse
	mapErr := mapstructure.Decode(apiRequest, &createVolume)
	if mapErr != nil {

		return nil, mapErr
	}

	return &createVolume, nil
}

// UpdateVolume updates the configuration of a Volume on the Silk server.
//
// Valid keys for the config are: `name`, `size`, `description`, `volume_group`, and `read_only`.
func (c *Credentials) UpdateVolume(name string, config map[string]interface{}, timeout ...int) (*CreateOrUpdateVolumeResponse, error) {

	httpTimeout := httpTimeout(timeout)

	// Validate that the user provided keys are valid for this API
	validUpdateKeys := []string{"name", "size", "description", "volume_group", "read_only"}
	var invalidUserProvidedKeys []string
	for key := range config {

		if c.stringInSlice(validUpdateKeys, key) == false {
			invalidUserProvidedKeys = append(invalidUserProvidedKeys, key)
		}
	}

	// Return an error message if any invalid keys are found
	if len(invalidUserProvidedKeys) != 0 {
		return nil, fmt.Errorf("The provided 'config' parameter contains invalid keys. 'name', 'size', 'description', 'volume_group', and 'read_only' are the only valid choices")
	}

	volumeID, err := c.GetVolumeID(name)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Patch(fmt.Sprintf("/volumes/%d", volumeID), config, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateOrUpdateVolumeResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// DeleteVolume deletes a Volume from the Silk server.
func (c *Credentials) DeleteVolume(name string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	volumeID, err := c.GetVolumeID(name)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Delete(fmt.Sprintf("/volumes/%d", volumeID), httpTimeout)
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

// GetVolumeID provides the ID for the provided host Volume name.
func (c *Credentials) GetVolumeID(name string, timeout ...int) (int, error) {

	httpTimeout := httpTimeout(timeout)

	objectsOnServer, err := c.GetVolumes(httpTimeout)
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
		return 0, fmt.Errorf("The server does not contain a Volume named '%s'", name)
	}

	return objectID, nil

}
