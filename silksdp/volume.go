package silksdp

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

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

// pre v1.4.0, full pull + client side filter. kept for reference
// // GetVolumes returns information on all Volumes found on the Silk server.
// func (c *Credentials) GetVolumes(timeout ...int) (*GetVolumesResponse, error) {
//
// 	httpTimeout := httpTimeout(timeout)
//
// 	apiRequest, err := c.Get("/volumes", httpTimeout)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Convert the API Response (map[string]interface{}) to a struct
// 	var apiResponse GetVolumesResponse
// 	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
// 	if mapErr != nil {
// 		return nil, mapErr
// 	}
//
// 	return &apiResponse, nil
// }

// GetVolumes returns information on all Volumes found on the Silk server.
func (c *Credentials) GetVolumes(timeout ...int) (*GetVolumesResponse, error) {
	var apiResponse GetVolumesResponse
	if err := c.queryInto("/volumes", nil, &apiResponse, httpTimeout(timeout)); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

// pre v1.4.0, full pull + client side filter. kept for reference
// func (c *Credentials) GetVolumeName(id int, timeout ...int) (*GetVolumesResponse, error) {
//
// 	httpTimeout := httpTimeout(timeout)
//
// 	apiRequest, err := c.Get(fmt.Sprintf("/volumes?id__in=%v", id), httpTimeout)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Convert the API Response (map[string]interface{}) to a struct
// 	var apiResponse GetVolumesResponse
// 	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
// 	if mapErr != nil {
// 		return nil, mapErr
// 	}
//
// 	return &apiResponse, nil
// }

// GetVolumeName returns the volume with the given id. Same as GetVolumeByID, kept for older callers.
func (c *Credentials) GetVolumeName(id int, timeout ...int) (*GetVolumesResponse, error) {
	return c.GetVolumeByID(id, timeout...)
}

// GetVolumeByID returns the volume with the given id (empty Hits if it doesnt exist).
func (c *Credentials) GetVolumeByID(id int, timeout ...int) (*GetVolumesResponse, error) {
	var apiResponse GetVolumesResponse
	if err := c.queryInto("/volumes", url.Values{"id": {strconv.Itoa(id)}}, &apiResponse, httpTimeout(timeout)); err != nil {
		return nil, err
	}
	return &apiResponse, nil
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

	// Remove Host mappings before remove volume
	hostMappings, err := c.GetVolumeHostMappings(name, httpTimeout)
	if len(hostMappings) > 0 {
		for _, hostName := range hostMappings {
			c.DeleteHostVolumeMapping(hostName, name, httpTimeout)
		}
	}
	// Remove Host group mappings before remove volume
	hostGroupMappings, err := c.GetVolumeHostGroupMappings(name, httpTimeout)
	if len(hostGroupMappings) > 0 {
		for _, hostGroupName := range hostGroupMappings {
			c.DeleteHostGroupVolumeMapping(hostGroupName, name, httpTimeout)
		}
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

// pre v1.4.0, full pull + client side filter. kept for reference
// // GetVolumeID provides the ID for the provided host Volume name.
// func (c *Credentials) GetVolumeID(name string, timeout ...int) (int, error) {
//
// 	httpTimeout := httpTimeout(timeout)
//
// 	volumes, err := c.GetVolumes(httpTimeout)
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	// Set volumeID to a value (-1) that can not be returned by the server
// 	volumeID := -1
// 	for _, volume := range volumes.Hits {
// 		if volume.Name == name {
// 			volumeID = volume.ID
// 		}
// 	}
//
// 	// If the volumeID has not been updated (i.e not found on the server) return an error message
// 	if volumeID == -1 {
// 		return 0, fmt.Errorf("The server does not contain a Volume named '%s'", name)
// 	}
//
// 	return volumeID, nil
// }

// GetVolumeID provides the ID for the provided host Volume name.
func (c *Credentials) GetVolumeID(name string, timeout ...int) (int, error) {
	volumes, err := c.GetVolumeByName(name, timeout...)
	if err != nil {
		return 0, err
	}
	// api name match ignores case, keep ours exact
	for _, volume := range volumes.Hits {
		if volume.Name == name {
			return volume.ID, nil
		}
	}
	return 0, fmt.Errorf("The server does not contain a Volume named '%s'", name)
}

// pre v1.4.0, full pull + client side filter. kept for reference
// // GetVolumeHostMappings returns all Hosts that are mapped to the provided Volume.
// func (c *Credentials) GetVolumeHostMappings(volumeName string, timeout ...int) ([]string, error) {
//
// 	httpTimeout := httpTimeout(timeout)
//
// 	volumeID, err := c.GetVolumeID(volumeName)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	hostMappingsOnServer, err := c.GetHostMappings(httpTimeout)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Filter out the user provided volume and host from the hostMappingsOnServer
// 	// results
// 	hostName := []string{}
//
// 	for _, mapping := range hostMappingsOnServer {
// 		if mapping.Volume.Ref == fmt.Sprintf("/volumes/%d", volumeID) {
// 			var hostID int
// 			if strings.Contains(mapping.Host.Ref, "/hosts") == true {
//
// 				hostRefID := strings.Replace(mapping.Host.Ref, "/hosts/", "", 1)
// 				hostID, err = strconv.Atoi(hostRefID)
// 				if err != nil {
// 					return nil, err
// 				}
//
// 				name, err := c.GetHostName(hostID)
// 				if err != nil {
// 					return nil, err
// 				}
//
// 				hostName = append(hostName, name)
// 			}
//
// 		}
//
// 	}
//
// 	// If the mappingID has not been updated (i.e not found on the server) return an error message
// 	// if len(hostName) == 0 {
// 	// 	return nil, fmt.Errorf("No Host Mappings found on the Volume '%s'", volumeName)
// 	// }
//
// 	return hostName, nil
// }

// GetVolumeHostMappings returns all Hosts that are mapped to the provided Volume.
func (c *Credentials) GetVolumeHostMappings(volumeName string, timeout ...int) ([]string, error) {
	volumeID, err := c.GetVolumeID(volumeName, timeout...)
	if err != nil {
		return nil, err
	}
	hosts, _, err := c.GetVolumeMappingsByID(volumeID, timeout...)
	return hosts, err
}

// GetVolumeMappingsByID returns the host names and host group names mapped to a volume, from one /mappings call
func (c *Credentials) GetVolumeMappingsByID(volumeID int, timeout ...int) ([]string, []string, error) {
	return c.mappedNames(ref("volumes", volumeID), httpTimeout(timeout))
}

// mappedNames turns the mappings on a volume (or volume group) ref into host and host group names.
// one id__in lookup per type instead of one per mapping
func (c *Credentials) mappedNames(volumeRef string, timeout int) ([]string, []string, error) {
	mappings, err := c.getMappings(url.Values{"volume.ref": {volumeRef}}, timeout)
	if err != nil {
		return nil, nil, err
	}

	var hostIDs, hostGroupIDs []int
	for _, m := range mappings {
		switch {
		case strings.HasPrefix(m.Host.Ref, "/hosts/"):
			hostIDs = append(hostIDs, refID(m.Host.Ref))
		case strings.HasPrefix(m.Host.Ref, "/host_groups/"):
			hostGroupIDs = append(hostGroupIDs, refID(m.Host.Ref))
		}
	}

	hostNames := []string{}
	if len(hostIDs) > 0 {
		var hosts GetHostsResponse
		if err := c.queryInto("/hosts", url.Values{"id__in": {idList(hostIDs)}}, &hosts, timeout); err != nil {
			return nil, nil, err
		}
		for _, h := range hosts.Hits {
			hostNames = append(hostNames, h.Name)
		}
	}

	hostGroupNames := []string{}
	if len(hostGroupIDs) > 0 {
		var hostGroups GetHostGroupsResponse
		if err := c.queryInto("/host_groups", url.Values{"id__in": {idList(hostGroupIDs)}}, &hostGroups, timeout); err != nil {
			return nil, nil, err
		}
		for _, hg := range hostGroups.Hits {
			hostGroupNames = append(hostGroupNames, hg.Name)
		}
	}

	return hostNames, hostGroupNames, nil
}

// pre v1.4.0, full pull + client side filter. kept for reference
// // GetVolumeHostGroupMappings returns all Host Groups that are mapped to the provided Volume.
// func (c *Credentials) GetVolumeHostGroupMappings(volumeName string, timeout ...int) ([]string, error) {
//
// 	httpTimeout := httpTimeout(timeout)
//
// 	volumeID, err := c.GetVolumeID(volumeName)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	hostGroupMappingsOnServer, err := c.GetHostMappings(httpTimeout)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Filter out the user provided volume and host from the hostMappingsOnServer
// 	// results
// 	hostName := []string{}
//
// 	for _, mapping := range hostGroupMappingsOnServer {
// 		if mapping.Volume.Ref == fmt.Sprintf("/volumes/%d", volumeID) {
// 			var hostGroupID int
// 			if strings.Contains(mapping.Host.Ref, "/host_groups") == true {
//
// 				hostGroupRefID := strings.Replace(mapping.Host.Ref, "/host_groups/", "", 1)
// 				hostGroupID, err = strconv.Atoi(hostGroupRefID)
// 				if err != nil {
// 					return nil, err
// 				}
//
// 				name, err := c.GetHostGroupName(hostGroupID)
// 				if err != nil {
// 					return nil, err
// 				}
//
// 				hostName = append(hostName, name)
// 			}
//
// 		}
//
// 	}
//
// 	// If the mappingID has not been updated (i.e not found on the server) return an error message
// 	// if len(hostName) == 0 {
// 	// 	return nil, fmt.Errorf("No Host Mappings found on the Volume '%s'", volumeName)
// 	// }
//
// 	return hostName, nil
// }

// GetVolumeHostGroupMappings returns all Host Groups that are mapped to the provided Volume.
func (c *Credentials) GetVolumeHostGroupMappings(volumeName string, timeout ...int) ([]string, error) {
	volumeID, err := c.GetVolumeID(volumeName, timeout...)
	if err != nil {
		return nil, err
	}
	_, hostGroups, err := c.GetVolumeMappingsByID(volumeID, timeout...)
	return hostGroups, err
}

// pre v1.4.0, full pull + client side filter. kept for reference
// // GetVolumeGroupHostGroupMappings returns all Host Groups that are mapped to the provided Volume Group.
// func (c *Credentials) GetVolumeGroupHostGroupMappings(volumeGroupName string, timeout ...int) ([]string, error) {
//
// 	httpTimeout := httpTimeout(timeout)
//
// 	volumeGroupID, err := c.GetVolumeGroupID(volumeGroupName)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	hostGroupMappingsOnServer, err := c.GetHostMappings(httpTimeout)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Filter out the user provided volume and host from the hostMappingsOnServer
// 	// results
// 	hostName := []string{}
//
// 	for _, mapping := range hostGroupMappingsOnServer {
// 		if mapping.Volume.Ref == fmt.Sprintf("/volume_groups/%d", volumeGroupID) {
// 			var hostGroupID int
// 			if strings.Contains(mapping.Host.Ref, "/host_groups") == true {
//
// 				hostGroupRefID := strings.Replace(mapping.Host.Ref, "/host_groups/", "", 1)
// 				hostGroupID, err = strconv.Atoi(hostGroupRefID)
// 				if err != nil {
// 					return nil, err
// 				}
//
// 				name, err := c.GetHostGroupName(hostGroupID)
// 				if err != nil {
// 					return nil, err
// 				}
//
// 				hostName = append(hostName, name)
// 			}
//
// 		}
//
// 	}
//
// 	// If the mappingID has not been updated (i.e not found on the server) return an error message
// 	// if len(hostName) == 0 {
// 	// 	return nil, fmt.Errorf("No Host Mappings found on the Volume Group '%s'", volumeGroupName)
// 	// }
//
// 	return hostName, nil
// }

// GetVolumeGroupHostGroupMappings returns all Host Groups that are mapped to the provided Volume Group.
func (c *Credentials) GetVolumeGroupHostGroupMappings(volumeGroupName string, timeout ...int) ([]string, error) {
	volumeGroupID, err := c.GetVolumeGroupID(volumeGroupName, timeout...)
	if err != nil {
		return nil, err
	}
	_, hostGroups, err := c.mappedNames(ref("volume_groups", volumeGroupID), httpTimeout(timeout))
	return hostGroups, err
}

// pre v1.4.0, full pull + client side filter. kept for reference
// // GetVolumeByName submits a strict API query for host objects of a specific name.
// func (c *Credentials) GetVolumeByName(volumename string, timeout ...int) (*GetVolumesResponse, error) {
//
// 	httpTimeout := httpTimeout(timeout)
//
// 	enduri := ("/volumes?name__contains=" + volumename)
//
// 	apiRequest, err := c.Get(enduri, httpTimeout)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Convert the API Response (map[string]interface{}) to a struct
// 	var apiResponse GetVolumesResponse
// 	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
// 	if mapErr != nil {
// 		return nil, mapErr
// 	}
//
// 	return &apiResponse, nil
// }

// GetVolumeByName returns volumes matching the name exactly (api match ignores case).
func (c *Credentials) GetVolumeByName(volumename string, timeout ...int) (*GetVolumesResponse, error) {
	var apiResponse GetVolumesResponse
	if err := c.queryInto("/volumes", url.Values{"name": {volumename}}, &apiResponse, httpTimeout(timeout)); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

