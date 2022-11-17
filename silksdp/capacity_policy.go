package silksdp

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// GetCapacityPolicy returns information on all Capacity Policys found on the Silk server.
func (c *Credentials) GetCapacityPolicy(timeout ...int) (*GetCapacityPolicyResponse, error) {
	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.Get("/vg_capacity_policies", httpTimeout)
	if err != nil {
		return nil, err
	}
	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse GetCapacityPolicyResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// GetCapacityPolicyID collects the capacity policy ID
func (c *Credentials) GetCapacityPolicyID(name string, timeout ...int) (int, error) {

	httpTimeout := httpTimeout(timeout)

	objectsOnServer, err := c.GetCapacityPolicy(httpTimeout)
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
		return 0, fmt.Errorf("The server does not contain a Capacity Policy named '%s'", name)
	}

	return objectID, nil

}

// CreateCapacityPolicy creates a new Capacity Policy on the Silk server.
func (c *Credentials) CreateCapacityPolicy(name string, warningthreshold int, errorthreshold int, criticalthreshold int, fullthreshold int, snapshotoverheadthreshold int, timeout ...int) (*CreateOrUpdateCapacityPolicyResponse, error) {

	httpTimeout := httpTimeout(timeout)

	config := map[string]interface{}{}
	config["name"] = name
	config["warning_threshold"] = warningthreshold
	config["error_threshold"] = errorthreshold
	config["critical_threshold"] = criticalthreshold
	config["full_threshold"] = fullthreshold
	config["snapshot_overhead_threshold"] = snapshotoverheadthreshold

	apiRequest, err := c.Post("/vg_capacity_policies", config, httpTimeout)
	if err != nil {
		return nil, err
	}

	var apiResponse CreateOrUpdateCapacityPolicyResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// UpdateCapacityPolicy updates the Capacity Policy with the provided config options.
//
// Valid config keys are: "name", "warningthreshold", "errorthreshold", "criticalthreshold", "fullthreshold", "snapshotoverheadthreshold".
func (c *Credentials) UpdateCapacityPolicy(name string, config map[string]interface{}, timeout ...int) (*CreateOrUpdateCapacityPolicyResponse, error) {
	httpTimeout := httpTimeout(timeout)

	// Validate that the user provided keys are valid for this API
	validUpdateKeys := []string{"name", "warningthreshold", "errorthreshold", "criticalthreshold", "fullthreshold", "snapshotoverheadthreshold"}
	var invalidUserProvidedKeys []string
	for key := range config {

		if c.stringInSlice(validUpdateKeys, key) == false {
			invalidUserProvidedKeys = append(invalidUserProvidedKeys, key)
		}
	}

	// Return an error message if any invalid keys are found
	if len(invalidUserProvidedKeys) != 0 {
		return nil, fmt.Errorf("The provided 'config' parameter contains invalid keys. 'name', 'warningthreshold', 'errorthreshold', 'criticalthreshold', 'fullthreshold', 'snapshotoverheadthreshold' are the only valid choices")
	}

	CapacityPolicyID, err := c.GetCapacityPolicyID(name, httpTimeout)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Patch(fmt.Sprintf("/vg_capacity_policies/%d", CapacityPolicyID), config, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateOrUpdateCapacityPolicyResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil

}

// DeleteCapacityPolicy deletes a Capacity Policy from the Silk server.
func (c *Credentials) DeleteCapacityPolicy(name string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	CapacityPolicyID, err := c.GetCapacityPolicyID(name, httpTimeout)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Delete(fmt.Sprintf("/vg_capacity_policies/%d", CapacityPolicyID), httpTimeout)
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
// GetCapacityPolicyByName returns information on all Capacity Policys found on the Silk server.
func (c *Credentials) GetCapacityPolicyByName(capacitypolicyname string, timeout ...int) (*GetCapacityPolicyResponse, error) {
	httpTimeout := httpTimeout(timeout)

	enduri := ("/vg_capacity_policies?name__contains=" + capacitypolicyname)

	apiRequest, err := c.Get(enduri, httpTimeout)
	if err != nil {
		return nil, err
	}
	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse GetCapacityPolicyResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}