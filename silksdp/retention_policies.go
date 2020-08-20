package silksdp

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// GetRetentionPolicy returns information on all Retention Policies found on the Silk server.
func (c *Credentials) GetRetentionPolicy(timeout ...int) (*GetRetentionPolicyResponse, error) {
	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.Get("/retention_policies", httpTimeout)
	if err != nil {
		return nil, err
	}
	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse GetRetentionPolicyResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// DeleteRetentionPolicy deletes a Retention Policy from the Silk server.
func (c *Credentials) DeleteRetentionPolicy(name string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	RetentionPolicyID, err := c.GetRetentionPolicyID(name, httpTimeout)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Delete(fmt.Sprintf("/retention_policies/%d", RetentionPolicyID), httpTimeout)
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

// GetRetentionPolicyID is a quick function for grabbing a retention policy object ID
func (c *Credentials) GetRetentionPolicyID(name string, timeout ...int) (int, error) {

	httpTimeout := httpTimeout(timeout)

	objectsOnServer, err := c.GetRetentionPolicy(httpTimeout)
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
		return 0, fmt.Errorf("The server does not contain a Retention Policy named '%s'", name)
	}

	return objectID, nil

}

// CreateRetentionPolicy creates a new Retention Policy on the Silk server.
func (c *Credentials) CreateRetentionPolicy(name string, numsnapshots string, weeks string, days string, hours string, timeout ...int) (*CreateOrUpdateRetentionPolicyResponse, error) {

	httpTimeout := httpTimeout(timeout)

	config := map[string]interface{}{}
	config["name"] = name
	config["num_snapshots"] = numsnapshots
	config["weeks"] = weeks
	config["days"] = days
	config["hours"] = hours

	apiRequest, err := c.Post("/retention_policies", config, httpTimeout)
	if err != nil {
		return nil, err
	}

	var apiResponse CreateOrUpdateRetentionPolicyResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// UpdateRetentionPolicy updates the Retention Policy with the provided config options.
//
// Valid config keys are: name, num_snapshots, weeks, days, and hours.
func (c *Credentials) UpdateRetentionPolicy(name string, config map[string]interface{}, timeout ...int) (*CreateOrUpdateRetentionPolicyResponse, error) {
	httpTimeout := httpTimeout(timeout)

	// Validate that the user provided keys are valid for this API
	validUpdateKeys := []string{"name", "num_snapshots", "weeks", "days", "hours"}
	var invalidUserProvidedKeys []string
	for key := range config {

		if c.stringInSlice(validUpdateKeys, key) == false {
			invalidUserProvidedKeys = append(invalidUserProvidedKeys, key)
		}
	}

	// Return an error message if any invalid keys are found
	if len(invalidUserProvidedKeys) != 0 {
		return nil, fmt.Errorf("The provided 'config' parameter contains invalid keys. 'name', 'num_snapshots', 'weeks', 'days' and 'hours' are the only valid choices")
	}

	RetentionPolicyID, err := c.GetRetentionPolicyID(name, httpTimeout)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Patch(fmt.Sprintf("/retention_policies/%d", RetentionPolicyID), config, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateOrUpdateRetentionPolicyResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil

}
