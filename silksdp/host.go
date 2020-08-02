package silksdp

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// CreateHost creates a new Host on the Silk server.
//
// Valid hostType choices are 'Linux', 'Windows', and 'ESX'.
func (c *Credentials) CreateHost(name, hostType string, timeout ...int) (*CreateOrUpdateHostResponse, error) {

	httpTimeout := httpTimeout(timeout)

	// Validate that the user provided hostTypes are valid
	validHostTypes := []string{"Linux", "Windows", "ESX"}

	if c.stringInSlice(validHostTypes, hostType) == false {
		return nil, fmt.Errorf("'%s' is not a valid hostType. Valid choices are 'Linux', 'Windows', and 'ESX'", hostType)
	}

	config := map[string]interface{}{}
	config["name"] = name
	config["type"] = hostType

	apiRequest, err := c.Post("/hosts", config, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateOrUpdateHostResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// GetHosts returns information on all Hosts found on the Silk server.
func (c *Credentials) GetHosts(timeout ...int) (*GetHostsResponse, error) {

	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.Get("/hosts", httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse GetHostsResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// UpdateHost updates the Host with the provided config options.
//
// Valid keys for the config map[string]interface{} are: name, type. and host_group.
func (c *Credentials) UpdateHost(name string, config map[string]interface{}, timeout ...int) (*CreateOrUpdateHostResponse, error) {
	httpTimeout := httpTimeout(timeout)

	// Validate that the user provided keys are valid for this API
	validUpdateKeys := []string{"name", "type", "host_group"}
	var invalidUserProvidedKeys []string
	for key := range config {

		if c.stringInSlice(validUpdateKeys, key) == false {
			invalidUserProvidedKeys = append(invalidUserProvidedKeys, key)
		}
	}

	// Return an error message if any invalid keys are found
	if len(invalidUserProvidedKeys) != 0 {
		return nil, fmt.Errorf("The provided 'config' parameter contains invalid keys. 'name' and 'type' are the only valid choices")
	}

	hostID, err := c.GetHostID(name)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Get(fmt.Sprintf("/hosts/%d", hostID), httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateOrUpdateHostResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil

}

// DeleteHost deletes a Host from the Silk server.
func (c *Credentials) DeleteHost(name string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostID, err := c.GetHostID(name)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Delete(fmt.Sprintf("/hosts/%d", hostID), httpTimeout)
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

// CreateHostVolumeMapping will map a Host to the provided Volume.
func (c *Credentials) CreateHostVolumeMapping(hostName, volumeName string, timeout ...int) (*CreateHostVolumeMappingResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostsOnServer, err := c.GetHosts(httpTimeout)
	if err != nil {
		return nil, err
	}

	// Validates that the provided host is not part of a Host Group which would prevent the host being added.
	for _, host := range hostsOnServer.Hits {
		if host.Name == hostName {
			if host.IsPartOfGroup == true {
				return nil, fmt.Errorf("Host '%s' is a member of a Host Group and can not individually be mapped to a volume", hostName)
			}
		}

	}

	hostID, err := c.GetHostID(hostName)
	if err != nil {
		return nil, err
	}

	volumeID, err := c.GetVolumeID(volumeName)
	if err != nil {
		return nil, err
	}

	hostConfig := map[string]interface{}{}
	hostConfig["ref"] = fmt.Sprintf("/hosts/%d", hostID)

	volumeConfig := map[string]string{}
	volumeConfig["ref"] = fmt.Sprintf("/volumes/%d", volumeID)

	config := map[string]interface{}{}
	config["host"] = hostConfig
	config["volume"] = volumeConfig

	apiRequest, err := c.Post("/mappings", config, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateHostVolumeMappingResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil

}

// GetHostMappings returns information on all Host Mappings found on the Silk server.
//
// The returned []HostMappingRespons slice only contains information on the hosts and not
// the full response of the API call. If no host mappings are found, an empty slice will be returned.
func (c *Credentials) GetHostMappings(timeout ...int) ([]IndividualHostMappingResponse, error) {

	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.Get("/mappings", httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse GetHostMappingsResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	// Filter all "hosts" mappings from the apiRequest and save them to a
	// slice that will be returned to the user.
	var hostMappings []IndividualHostMappingResponse
	for _, value := range apiResponse.Hits {
		if strings.Contains(value.Host.Ref, "/hosts") == true {
			hostMappings = append(hostMappings, value)
		}
	}

	return hostMappings, nil
}

// DeleteHostMappings removes all mappings from the provided host.
func (c *Credentials) DeleteHostMappings(hostName string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostID, err := c.GetHostID(hostName)
	if err != nil {
		return nil, err
	}

	hostMappingsOnServer, err := c.GetHostMappings(httpTimeout)
	if err != nil {
		return nil, err
	}

	// Filter all "hosts" mappings found on the server and save to a new
	// slice for processing
	var mappingIDs []int
	for _, mapping := range hostMappingsOnServer {
		if mapping.Host.Ref == fmt.Sprintf("/hosts/%d", hostID) {
			mappingIDs = append(mappingIDs, mapping.ID)
		}

	}

	// Return an error message if the host does not have any mappings
	if len(mappingIDs) == 0 {
		return nil, fmt.Errorf("No mappings found on the host '%s'", hostName)
	}

	// Loop through every mapping id in the mappingIDs slice and execute a delete call on that
	// id
	for _, id := range mappingIDs {
		_, err := c.Delete(fmt.Sprintf("/mappings/%d", id), httpTimeout)
		if err != nil {
			return nil, err
		}

	}

	// Since we are ignoring the response of each of the Delete calls above,
	// create a "dummy" DeleteReponse to return to the end user to signify success
	var apiResponse DeleteResponse
	apiResponse.StatusCode = 204

	return &apiResponse, nil
}

// DeleteHostVolumeMapping removes a single Volume Mapping from a Host.
func (c *Credentials) DeleteHostVolumeMapping(hostName, volumeName string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostID, err := c.GetHostID(hostName)
	if err != nil {
		return nil, err
	}

	volumeID, err := c.GetVolumeID(volumeName)
	if err != nil {
		return nil, err
	}

	hostMappingsOnServer, err := c.GetHostMappings(httpTimeout)
	if err != nil {
		return nil, err
	}

	// Filter out the user provided volume and host from the hostMappingsOnServer
	// results
	mappingID := -1
	for _, mapping := range hostMappingsOnServer {
		if mapping.Volume.Ref == fmt.Sprintf("/volumes/%d", volumeID) {
			if mapping.Host.Ref == fmt.Sprintf("/hosts/%d", hostID) {
				mappingID = mapping.ID

			}
		}

	}

	// If the mappingID has not been updated (i.e not found on the server) return an error message
	if mappingID == -1 {
		return nil, fmt.Errorf("No %s Volume Mappings found on the Host '%s'", volumeName, hostName)
	}

	apiRequest, err := c.Delete(fmt.Sprintf("/mappings/%d", mappingID), httpTimeout)
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

// GetHostID provides the ID for the provided Host name.
func (c *Credentials) GetHostID(name string, timeout ...int) (int, error) {

	httpTimeout := httpTimeout(timeout)

	objectsOnServer, err := c.GetHosts(httpTimeout)
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
		return 0, fmt.Errorf("The server does not contain a Host named '%s'", name)
	}

	return objectID, nil

}

// CreateHostPWWN adds a PWWN to a Host.
func (c *Credentials) CreateHostPWWN(hostName, PWWN string, timeout ...int) (*CreateHostPWWNResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostID, err := c.GetHostID(hostName)
	if err != nil {
		return nil, err
	}

	hostConfig := map[string]interface{}{}
	hostConfig["ref"] = fmt.Sprintf("/hosts/%d", hostID)

	config := map[string]interface{}{}
	config["pwwn"] = PWWN
	config["host"] = hostConfig

	apiRequest, err := c.Post("/host_fc_ports", config, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateHostPWWNResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// GetHostPWWN returns all PWWNs that have been added to the Host.
//
// The returned []IndividualHostPWWNResponse slice only contains information on the Host PWWN mappings and not
// the full response of the API call. If no PWWNs have been added to a Host, an empty slice will be returned.
func (c *Credentials) GetHostPWWN(hostName string, timeout ...int) ([]IndividualHostPWWNResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostID, err := c.GetHostID(hostName)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Get("/host_fc_ports", httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse GetHostPWWNResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	// Filter out all host mappings from the apiRequest
	var hostPWWN []IndividualHostPWWNResponse
	for _, value := range apiResponse.Hits {
		if value.Host.Ref == fmt.Sprintf("/hosts/%d", hostID) {
			hostPWWN = append(hostPWWN, value)

		}
	}

	return hostPWWN, nil
}

// DeleteHostPWWN removes all PWWNs from a Host.
func (c *Credentials) DeleteHostPWWN(hostName string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostPWWNs, err := c.GetHostPWWN(hostName)
	if err != nil {
		return nil, err
	}

	// Get the host id of each PWWN/Host mapping returned in hostPWWNs
	var pwwnToDelete []int
	for _, host := range hostPWWNs {
		pwwnToDelete = append(pwwnToDelete, host.ID)
	}

	// If the pwwnToDelete slice is empty (i.e no PWWN mappings found for the host) return an error message
	if len(pwwnToDelete) == 0 {
		return nil, fmt.Errorf("No PWWNs found on the host '%s'", hostName)
	}

	// Loop through every id in the pwwnToDelete slice and execute a delete call on that
	// id
	for _, id := range pwwnToDelete {
		_, err := c.Delete(fmt.Sprintf("/host_fc_ports/%d", id), httpTimeout)
		if err != nil {
			return nil, err
		}

	}

	// Since we are ignoring the response of each of the Delete calls above,
	// create a "dummy" DeleteReponse to return to the end user to signify success
	var apiResponse DeleteResponse
	apiResponse.StatusCode = 204

	return &apiResponse, nil
}

// GetHostIQN returns all IQNs that have been added to the Host.
//
// The returned []IndividualHostIQNResponse slice only contains information on the Host IQN mappings and not
// the full response of the API call. If no IQNs have been added to a Host, an empty slice will be returned.
func (c *Credentials) GetHostIQN(hostName string, timeout ...int) ([]IndividualHostIQNResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostID, err := c.GetHostID(hostName)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Get("/host_iqns", httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse GetHostIQNResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	var hostIQN []IndividualHostIQNResponse
	for _, value := range apiResponse.Hits {

		if value.Host.Ref == fmt.Sprintf("/hosts/%d", hostID) {
			hostIQN = append(hostIQN, value)

		}
	}

	return hostIQN, nil
}

// DeleteHostIQN removes all IQNs from a Host.
func (c *Credentials) DeleteHostIQN(hostName string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostIQNs, err := c.GetHostIQN(hostName)
	if err != nil {
		return nil, err
	}

	// Get the host id of each IQN/Host mapping returned in hostIQNs
	var iqnToDelete []int
	for _, host := range hostIQNs {
		iqnToDelete = append(iqnToDelete, host.ID)

	}

	// If the iqnToDelete slice is empty (i.e no PWWN mappings found for the host) return an error message
	if len(iqnToDelete) == 0 {
		return nil, fmt.Errorf("No IQNs found on the host '%s'", hostName)
	}

	for _, id := range iqnToDelete {
		_, err := c.Delete(fmt.Sprintf("/host_iqns/%d", id), httpTimeout)
		if err != nil {
			return nil, err
		}

	}

	// Since we are ignoring the response of each of the Delete calls above,
	// create a "dummy" DeleteReponse to return to the end user to signify success
	var apiResponse DeleteResponse
	apiResponse.StatusCode = 204

	return &apiResponse, nil

}

// CreateHostIQN adds a IQN to a Host.
func (c *Credentials) CreateHostIQN(hostName, IQN string, timeout ...int) (*CreateHostIQNResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostID, err := c.GetHostID(hostName)
	if err != nil {
		return nil, err
	}

	hostConfig := map[string]interface{}{}
	hostConfig["ref"] = fmt.Sprintf("/hosts/%d", hostID)

	config := map[string]interface{}{}
	config["iqn"] = IQN
	config["host"] = hostConfig

	apiRequest, err := c.Post("/host_iqns", config, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateHostIQNResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}
