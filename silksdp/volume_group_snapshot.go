package silksdp

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// GetVolumeGroupSnapshot returns information on all Volume Group Snapshots found on the Silk server.
func (c *Credentials) GetVolumeGroupSnapshot(timeout ...int) (*GetVolumeGroupSnapshotResponse, error) { // <- here
	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.Get("/snapshots", httpTimeout) // <- here
	if err != nil {
		return nil, err
	}
	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse GetVolumeGroupSnapshotResponse // <- here
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// GetVolumeGroupSnapshotID helper function to get snapshot by ID
func (c *Credentials) GetVolumeGroupSnapshotID(name string, timeout ...int) (int, error) { // <- here

	httpTimeout := httpTimeout(timeout)

	objectsOnServer, err := c.GetVolumeGroupSnapshot(httpTimeout) // <- here
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
		return 0, fmt.Errorf("The server does not contain a Volume Group Snapshot named '%s'", name) // <- here
	}

	return objectID, nil

}

// CreateVolumeGroupSnapshot creates a new Volume Group Snapshot on the Silk server.
func (c *Credentials) CreateVolumeGroupSnapshot(name string, volumegroupname string, retentionpolicyname string, deletable bool, exposable bool, timeout ...int) (*CreateOrUpdateVolumeGroupSnapshotResponse, error) { // <- here

	httpTimeout := httpTimeout(timeout)

	// Get volume id from name and construct ref path
	volumeGroupID, err := c.GetVolumeGroupID(volumegroupname, httpTimeout)
	if err != nil {
		return nil, err
	}

	volumegrouppath := fmt.Sprintf("@{ref=/volume_groups/%d", volumeGroupID)

	// Get rep retention policy id from name and construct ref path
	retentionpolicyID, err := c.GetRetentionPolicyID(retentionpolicyname, httpTimeout)
	if err != nil {
		return nil, err
	}

	retentionpolicypath := fmt.Sprintf("@{ref=/retention_policies/%d", retentionpolicyID)

	config := map[string]interface{}{}
	config["name"] = name
	config["volume_group"] = volumegrouppath
	config["retention_policy"] = retentionpolicypath
	config["deletable"] = deletable
	config["exposable"] = exposable

	apiRequest, err := c.Post("/snapshots", config, httpTimeout) // <- here
	if err != nil {
		return nil, err
	}

	var apiResponse CreateOrUpdateVolumeGroupSnapshotResponse // <- here
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// UpdateVolumeGroupSnapshot updates the Volume Group Snapshot with the provided config options.
//
// Valid config keys are: name, num_snapshots, weeks, days, and hours.

/* UpdateVolumeGroupSnapshot not required as no PATCH support in the /snapshots endpoint
func (c *Credentials) UpdateVolumeGroupSnapshot(name string, config map[string]interface{}, timeout ...int) (*CreateOrUpdateVolumeGroupSnapshotResponse, error) { // <- here
	httpTimeout := httpTimeout(timeout)

	// Validate that the user provided keys are valid for this API
	validUpdateKeys := []string{"name", "volumegroup", "retentionpolicy"}
	var invalidUserProvidedKeys []string
	for key := range config {

		if c.stringInSlice(validUpdateKeys, key) == false {
			invalidUserProvidedKeys = append(invalidUserProvidedKeys, key)
		}
	}

	// Return an error message if any invalid keys are found
	if len(invalidUserProvidedKeys) != 0 {
		return nil, fmt.Errorf("The provided 'config' parameter contains invalid keys. 'name', 'volumegroup', 'retentionpolicy' are the only valid choices")
	}

	VolumeGroupSnapshotID, err := c.GetVolumeGroupSnapshotID(name, httpTimeout) // <- here
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Patch(fmt.Sprintf("/retention_policies/%d", VolumeGroupSnapshotID), config, httpTimeout) // <- here
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateOrUpdateVolumeGroupSnapshotResponse // <- here
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil

}
*/

// DeleteVolumeGroupSnapshot deletes a Volume Group Snapshot from the Silk server.
func (c *Credentials) DeleteVolumeGroupSnapshot(name string, timeout ...int) (*DeleteResponse, error) { // <- here

	httpTimeout := httpTimeout(timeout)

	VolumeGroupSnapshotID, err := c.GetVolumeGroupSnapshotID(name, httpTimeout) // <- here
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Delete(fmt.Sprintf("/snapshots/%d", VolumeGroupSnapshotID), httpTimeout) // <- here
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
