package silksdp

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// CreateHostGroup creates a new Host Group on the Silk server.
//
// allowDifferentHostTypes corresponds to the "Enable mixed host OS types" checkbox in the UI.
func (c *Credentials) CreateHostGroup(name, description string, allowDifferentHostTypes bool, timeout ...int) (*CreateOrUpdateHostGroupResponse, error) {

	httpTimeout := httpTimeout(timeout)

	config := map[string]interface{}{}
	config["name"] = name
	config["description"] = description
	config["allow_different_host_types"] = allowDifferentHostTypes

	apiRequest, err := c.Post("/host_groups", config, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateOrUpdateHostGroupResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// GetHostGroups returns information on all Host Groups found on the Silk server.
func (c *Credentials) GetHostGroups(timeout ...int) (*GetHostGroupsResponse, error) {

	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.Get("/host_groups", httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse GetHostGroupsResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}

// UpdateHostGroup updates the Host with the provided config options.
//
// Valid keys for the config map[string]interface{} are: description and allow_different_host_types.
// Valid config keys are:
func (c *Credentials) UpdateHostGroup(name string, config map[string]interface{}, timeout ...int) (*CreateOrUpdateHostGroupResponse, error) {
	httpTimeout := httpTimeout(timeout)

	// Validate that the user provided keys are valid for this API
	validUpdateKeys := []string{"description", "allow_different_host_types"}
	var invalidUserProvidedKeys []string
	for key := range config {

		if c.stringInSlice(validUpdateKeys, key) == false {
			invalidUserProvidedKeys = append(invalidUserProvidedKeys, key)
		}
	}

	// Return an error message if any invalid keys are found
	if len(invalidUserProvidedKeys) != 0 {
		return nil, fmt.Errorf("The provided 'config' parameter contains invalid keys. 'description' and 'allow_different_host_types' are the only valid choices")
	}

	hostGroupID, err := c.GetHostGroupID(name, httpTimeout)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Patch(fmt.Sprintf("/host_groups/%d", hostGroupID), config, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse CreateOrUpdateHostGroupResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil

}

// DeleteHostGroup deletes a Host Group from the Silk server.
func (c *Credentials) DeleteHostGroup(name string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostGroupID, err := c.GetHostGroupID(name, httpTimeout)
	if err != nil {
		return nil, err
	}

	apiRequest, err := c.Delete(fmt.Sprintf("/host_groups/%d", hostGroupID), httpTimeout)
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

// GetHostGroupID provides the ID for the provided Host Group name.
func (c *Credentials) GetHostGroupID(name string, timeout ...int) (int, error) {

	httpTimeout := httpTimeout(timeout)

	objectsOnServer, err := c.GetHostGroups(httpTimeout)
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
		return 0, fmt.Errorf("The server does not contain a Host Group named '%s'", name)
	}

	return objectID, nil

}

// GetHostGroupName provides the name of a Host Group given its ID.
func (c *Credentials) GetHostGroupName(id int, timeout ...int) (string, error) {

	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.Get(fmt.Sprintf("/host_groups?id__in=%d", id), httpTimeout)
	if err != nil {
		return "", err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse GetHostGroupsResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return "", mapErr
	}

	for _, hostGroup := range apiResponse.Hits {
		return hostGroup.Name, nil
	}

	return "", fmt.Errorf("Did not found hostgroup with id=%d", id)

}

// CreateHostGroupVolumeMapping will map a Host to the provided Volume.
func (c *Credentials) CreateHostGroupVolumeMapping(hostGroupName, volumeName string, timeout ...int) (*CreateHostVolumeMappingResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostGroupID, err := c.GetHostGroupID(hostGroupName, httpTimeout)
	if err != nil {
		return nil, err
	}

	volumeID, err := c.GetVolumeID(volumeName, httpTimeout)
	if err != nil {
		return nil, err
	}

	hostConfig := map[string]interface{}{}
	hostConfig["ref"] = fmt.Sprintf("/host_groups/%d", hostGroupID)

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

// // CreateHostGroupVolumeGroupMapping will map a Host Group to the provided Volume.
// func (c *Credentials) CreateHostGroupVolumeGroupMapping(hostGroupName string, volumeGroupName string, timeout ...int) (*CreateHostVolumeMappingResponse, error) {

// 	httpTimeout := httpTimeout(timeout)

// 	hostGroupID, err := c.GetHostGroupID(hostGroupName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	volumeGroupID, err := c.GetVolumeGroupID(volumeGroupName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	hostGroupConfig := map[string]interface{}{}
// 	hostGroupConfig["ref"] = fmt.Sprintf("/host_groups/%d", hostGroupID)

// 	volumeGroupConfig := map[string]string{}
// 	volumeGroupConfig["ref"] = fmt.Sprintf("/volume_groups/%d", volumeGroupID)

// 	config := map[string]interface{}{}
// 	config["host"] = hostGroupConfig
// 	config["volume"] = volumeGroupConfig

// 	apiRequest, err := c.Post("/mappings", config, httpTimeout)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Convert the API Response (map[string]interface{}) to a struct
// 	var apiResponse CreateHostVolumeMappingResponse
// 	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
// 	if mapErr != nil {
// 		return nil, mapErr
// 	}
// 	return &apiResponse, nil
// }

// CreateHostGroupVolumeGroupMapping will map a Host Group to the provided Volume.
// func (c *Credentials) CreateHostGroupVolumeGroupMapping(hostGroupName string, volumeGroupName string, timeout ...int) ([]CreateHostVolumeMappingResponse, error) {

// 	httpTimeout := httpTimeout(timeout)

// 	hostGroupID, err := c.GetHostGroupID(hostGroupName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var hostGroupVolumeMappingResponse []CreateHostVolumeMappingResponse

// 	volumesInVolumeGroup, err := c.GetVolumeGroupVolumes(volumeGroupName, httpTimeout)
// 	for _, volume := range volumesInVolumeGroup {
// 		hostGroupConfig := map[string]interface{}{}
// 		hostGroupConfig["ref"] = fmt.Sprintf("/host_groups/%d", hostGroupID)

// 		volumeGroupID, err := c.GetVolumeID(volume)
// 		if err != nil {
// 			return nil, err
// 		}
// 		volumeConfig := map[string]string{}
// 		volumeConfig["ref"] = fmt.Sprintf("/volumes/%d", volumeGroupID)

// 		config := map[string]interface{}{}
// 		config["host"] = hostGroupConfig
// 		config["volume"] = volumeConfig

// 		apiRequest, err := c.Post("/mappings", config, httpTimeout)
// 		if err != nil {
// 			return nil, err
// 		}

// 		var apiResponse CreateHostVolumeMappingResponse
// 		mapErr := mapstructure.Decode(apiRequest, &apiResponse)
// 		if mapErr != nil {
// 			return nil, mapErr
// 		}
// 		hostGroupVolumeMappingResponse = append(hostGroupVolumeMappingResponse,apiResponse)

// 	}

// 	return hostGroupVolumeMappingResponse, nil
// }

func (c *Credentials) CreateHostGroupVolumeGroupMapping(hostGroupName string, volumeGroupName string, timeout ...int) ([]CreateHostVolumeMappingResponse, error) {

	httpTimeout := httpTimeout(timeout)
	var hostGroupVolumeMappingResponse []CreateHostVolumeMappingResponse

	volumesInVolumeGroup, err := c.GetVolumeGroupVolumes(volumeGroupName, httpTimeout)
	if err != nil {
		return nil, err
	}
	for _, volume := range volumesInVolumeGroup {

		apiResponse, err := c.CreateHostGroupVolumeMapping(hostGroupName,volume,httpTimeout)
		if err != nil {
			return nil, err
		}

		hostGroupVolumeMappingResponse = append(hostGroupVolumeMappingResponse,*apiResponse)

	}
	return hostGroupVolumeMappingResponse, nil
}

// GetHostGroupMappings returns information on all Host Group Mappings found on the Silk server.
//
// The returned []HostMappingRespons slice only contains information on the Host Groups and not
// the full response of the API call. If no host mappings are found, an empty slice will be returned.
func (c *Credentials) GetHostGroupMappings(timeout ...int) ([]IndividualHostMappingResponse, error) {

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

	// Filter all "host_groups" mappings from the apiRequest and save them to a
	// slice that will be returned to the user.
	var hostGroupMappings []IndividualHostMappingResponse
	for _, value := range apiResponse.Hits {
		if strings.Contains(value.Host.Ref, "/host_groups") == true {
			hostGroupMappings = append(hostGroupMappings, value)
		}
	}

	return hostGroupMappings, nil

}

// DeleteHostGroupMappings removes all mappings from the provided Host Group.
func (c *Credentials) DeleteHostGroupMappings(hostGroupName string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostGroupID, err := c.GetHostGroupID(hostGroupName)
	if err != nil {
		return nil, err
	}

	hostGroupMappingsOnServer, err := c.GetHostGroupMappings(httpTimeout)
	if err != nil {
		return nil, err
	}

	// Filter all "host_groups" mappings found on the server and save to a new
	// slice for processing
	var mappingIDs []int
	for _, mapping := range hostGroupMappingsOnServer {
		if mapping.Host.Ref == fmt.Sprintf("/host_groups/%d", hostGroupID) {
			mappingIDs = append(mappingIDs, mapping.ID)
		}

	}

	// Return an error message if the host group does not have any mappings
	if len(mappingIDs) == 0 {
		return nil, fmt.Errorf("No mappings found on the Host Group '%s'", hostGroupName)
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

// DeleteHostGroupVolumeMapping removes a single Volume mapping from a Host Group.
func (c *Credentials) DeleteHostGroupVolumeMapping(hostGroupName, volumeName string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostGroupID, err := c.GetHostGroupID(hostGroupName)
	if err != nil {
		return nil, err
	}

	volumeID, err := c.GetVolumeID(volumeName)
	if err != nil {
		return nil, err
	}

	hostGroupMappingsOnServer, err := c.GetHostGroupMappings(httpTimeout)
	if err != nil {
		return nil, err
	}

	// Filter out the user provided volume and host group from the hostMappingsOnServer
	// results
	mappingID := -1
	for _, mapping := range hostGroupMappingsOnServer {
		if mapping.Volume.Ref == fmt.Sprintf("/volumes/%d", volumeID) {
			if mapping.Host.Ref == fmt.Sprintf("/host_groups/%d", hostGroupID) {
				mappingID = mapping.ID

			}
		}
	}

	// If the mappingID has not been updated (i.e not found on the server) return an error message
	if mappingID == -1 {
		return nil, fmt.Errorf("No %s Volume mappings found on the Host Group '%s'", volumeName, hostGroupName)
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

// DeleteHostGroupVolumeGroupMapping removes a single Volume Group mapping from a Host Group.
func (c *Credentials) DeleteHostGroupVolumeGroupMapping(hostGroupName, volumeGroupName string, timeout ...int) (*DeleteResponse, error) {

	httpTimeout := httpTimeout(timeout)

	hostGroupID, err := c.GetHostGroupID(hostGroupName)
	if err != nil {
		return nil, err
	}

	volumeGroupID, err := c.GetVolumeGroupID(volumeGroupName)
	if err != nil {
		return nil, err
	}

	hostGroupMappingsOnServer, err := c.GetHostGroupMappings(httpTimeout)
	if err != nil {
		return nil, err
	}

	// Filter out the user provided volume and host group from the hostMappingsOnServer
	// results
	mappingID := -1
	for _, mapping := range hostGroupMappingsOnServer {
		if mapping.Volume.Ref == fmt.Sprintf("/volume_groups/%d", volumeGroupID) {
			if mapping.Host.Ref == fmt.Sprintf("/host_groups/%d", hostGroupID) {
				mappingID = mapping.ID

			}
		}
	}

	// If the mappingID has not been updated (i.e not found on the server) return an error message
	if mappingID == -1 {
		return nil, fmt.Errorf("No %s Volume Group mappings found on the Host Group '%s'", volumeGroupName, hostGroupName)
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

// GetHostGroupHosts provides the name of each Host in a Host Group.
func (c *Credentials) GetHostGroupHosts(name string, timeout ...int) ([]string, error) {

	httpTimeout := httpTimeout(timeout)

	hostGroupID, err := c.GetHostGroupID(name, httpTimeout)
	if err != nil {
		return nil, err
	}

	hostsOnServer, err := c.GetHosts(httpTimeout)
	if err != nil {
		return nil, err
	}

	hostsInHostGroup := []string{}

	for _, host := range hostsOnServer.Hits {
		if host.HostGroup.Ref == fmt.Sprintf("/host_groups/%d", hostGroupID) {
			hostsInHostGroup = append(hostsInHostGroup, host.Name)
		}
	}

	return hostsInHostGroup, nil
}
