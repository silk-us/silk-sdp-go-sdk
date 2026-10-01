package silksdp

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

// pre v1.4.0, full pull + client side filter. kept for reference
// // GetVolumeGroupSnapshot returns information on all Volume Group Snapshots found on the Silk server.
// func (c *Credentials) GetVolumeGroupSnapshot(timeout ...int) (*GetVolumeGroupSnapshotResponse, error) { // <- here
// 	httpTimeout := httpTimeout(timeout)
//
// 	apiRequest, err := c.Get("/snapshots", httpTimeout) // <- here
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Convert the API Response (map[string]interface{}) to a struct
// 	var apiResponse GetVolumeGroupSnapshotResponse // <- here
// 	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
// 	if mapErr != nil {
// 		return nil, mapErr
// 	}
//
// 	return &apiResponse, nil
// }

// GetVolumeGroupSnapshot returns information on all Volume Group Snapshots found on the Silk server.
func (c *Credentials) GetVolumeGroupSnapshot(timeout ...int) (*GetVolumeGroupSnapshotResponse, error) {
	var apiResponse GetVolumeGroupSnapshotResponse
	if err := c.queryInto("/snapshots", nil, &apiResponse, httpTimeout(timeout)); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

// pre v1.4.0, full pull + client side filter. kept for reference
// // GetVolumeGroupSnapshotID helper function to get snapshot by ID
// func (c *Credentials) GetVolumeGroupSnapshotID(name string, timeout ...int) (int, error) { // <- here
//
// 	httpTimeout := httpTimeout(timeout)
//
// 	objectsOnServer, err := c.GetVolumeGroupSnapshot(httpTimeout) // <- here
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	// Set objectID to a value (-1) that can not be returned by the server
// 	objectID := -1
// 	for _, object := range objectsOnServer.Hits {
// 		if object.Name == name {
// 			objectID = object.ID
// 		}
//
// 	}
//
// 	// If the objectID has not been updated (i.e not found on the server) return an error message
// 	if objectID == -1 {
// 		return 0, fmt.Errorf("The server does not contain a Volume Group Snapshot named '%s'", name) // <- here
// 	}
//
// 	return objectID, nil
//
// }

// GetVolumeGroupSnapshotID returns the id of a snapshot by its full name (vg:snap).
// A bare short name works too as long as only one volume group has a snapshot by that name.
func (c *Credentials) GetVolumeGroupSnapshotID(name string, timeout ...int) (int, error) {
	q := url.Values{"name": {name}}
	if !strings.Contains(name, ":") {
		q = url.Values{"short_name": {name}}
	}

	var snaps GetVolumeGroupSnapshotResponse
	if err := c.queryInto("/snapshots", q, &snaps, httpTimeout(timeout)); err != nil {
		return 0, err
	}

	switch len(snaps.Hits) {
	case 0:
		return 0, fmt.Errorf("The server does not contain a Volume Group Snapshot named '%s'", name)
	case 1:
		return snaps.Hits[0].ID, nil
	}
	return 0, fmt.Errorf("More than one Volume Group Snapshot is named '%s', use the full volumegroup:snapshot name", name)
}

// GetVolumeGroupSnapshotIDInGroup returns the id of a snapshot inside a known volume group.
// Takes either "vg:snap" or just "snap".
func (c *Credentials) GetVolumeGroupSnapshotIDInGroup(volumeGroupID int, name string, timeout ...int) (int, error) {
	shortName := name
	if i := strings.LastIndex(name, ":"); i >= 0 {
		shortName = name[i+1:]
	}

	q := url.Values{
		"volume_group.ref": {ref("volume_groups", volumeGroupID)},
		"short_name":       {shortName},
	}
	var snaps GetVolumeGroupSnapshotResponse
	if err := c.queryInto("/snapshots", q, &snaps, httpTimeout(timeout)); err != nil {
		return 0, err
	}
	if len(snaps.Hits) == 0 {
		return 0, fmt.Errorf("The server does not contain a Volume Group Snapshot named '%s'", name)
	}
	return snaps.Hits[0].ID, nil
}

// pre v1.4.0, kept for reference
// // CreateVolumeGroupSnapshot creates a new Volume Group Snapshot on the Silk server.
// func (c *Credentials) CreateVolumeGroupSnapshot(name string, volumegroupname string, retentionpolicyname string, deletable bool, exposable bool, timeout ...int) (*CreateOrUpdateVolumeGroupSnapshotResponse, error) { // <- here
//
// 	httpTimeout := httpTimeout(timeout)
//
// 	// Get volume id from name and construct ref path
// 	volumeGroupID, err := c.GetVolumeGroupID(volumegroupname, httpTimeout)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	volumegrouppath := fmt.Sprintf("@{ref=/volume_groups/%d", volumeGroupID)
//
// 	// Get rep retention policy id from name and construct ref path
// 	retentionpolicyID, err := c.GetRetentionPolicyID(retentionpolicyname, httpTimeout)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	retentionpolicypath := fmt.Sprintf("@{ref=/retention_policies/%d", retentionpolicyID)
//
// 	config := map[string]interface{}{}
// 	config["name"] = name
// 	config["volume_group"] = volumegrouppath
// 	config["retention_policy"] = retentionpolicypath
// 	config["deletable"] = deletable
// 	config["exposable"] = exposable
//
// 	apiRequest, err := c.Post("/snapshots", config, httpTimeout) // <- here
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var apiResponse CreateOrUpdateVolumeGroupSnapshotResponse // <- here
// 	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
// 	if mapErr != nil {
// 		return nil, mapErr
// 	}
//
// 	return &apiResponse, nil
// }

// CreateVolumeGroupSnapshot takes a snapshot of a Volume Group.
//
// `name` is the short name, the array makes the full name volumegroup:name.
// `deletable` sets is_auto_deleteable. Pass false for anything terraform owns or the
// retention policy can delete it out from under the state file.
// `exposable` should be false for a plain snapshot, use CreateVolumeGroupView for views.
func (c *Credentials) CreateVolumeGroupSnapshot(name string, volumegroupname string, retentionpolicyname string, deletable bool, exposable bool, timeout ...int) (*CreateOrUpdateVolumeGroupSnapshotResponse, error) {

	httpTimeout := httpTimeout(timeout)

	volumeGroupID, err := c.GetVolumeGroupID(volumegroupname, httpTimeout)
	if err != nil {
		return nil, err
	}

	retentionPolicyID, err := c.GetRetentionPolicyID(retentionpolicyname, httpTimeout)
	if err != nil {
		return nil, err
	}

	config := map[string]interface{}{}
	config["short_name"] = name
	config["source"] = map[string]string{"ref": ref("volume_groups", volumeGroupID)}
	config["retention_policy"] = map[string]string{"ref": ref("retention_policies", retentionPolicyID)}
	config["is_auto_deleteable"] = deletable
	config["is_exposable"] = exposable

	return c.postSnapshot(config, volumegroupname+":"+name, httpTimeout)
}

// CreateVolumeGroupView creates a writable, mappable view of an existing snapshot.
//
// `name` is the short name of the view. `snapshotName` is the snapshot to view, vg:snap or just snap.
// `deletable` sets is_auto_deleteable, same caveat as snapshots.
func (c *Credentials) CreateVolumeGroupView(name string, snapshotName string, retentionpolicyname string, deletable bool, timeout ...int) (*CreateOrUpdateVolumeGroupSnapshotResponse, error) {

	httpTimeout := httpTimeout(timeout)

	snapshotID, err := c.GetVolumeGroupSnapshotID(snapshotName, httpTimeout)
	if err != nil {
		return nil, err
	}

	// need the parent's vg prefix to find the view after
	parent, err := c.GetVolumeGroupSnapshotByID(snapshotID, httpTimeout)
	if err != nil {
		return nil, err
	}
	if len(parent.Hits) == 0 {
		return nil, fmt.Errorf("The server does not contain a Volume Group Snapshot named '%s'", snapshotName)
	}
	vgName := parent.Hits[0].Name[:strings.LastIndex(parent.Hits[0].Name, ":")]

	retentionPolicyID, err := c.GetRetentionPolicyID(retentionpolicyname, httpTimeout)
	if err != nil {
		return nil, err
	}

	config := map[string]interface{}{}
	config["short_name"] = name
	config["source"] = map[string]string{"ref": ref("snapshots", snapshotID)}
	config["retention_policy"] = map[string]string{"ref": ref("retention_policies", retentionPolicyID)}
	config["is_auto_deleteable"] = deletable
	config["is_exposable"] = true

	return c.postSnapshot(config, vgName+":"+name, httpTimeout)
}

// postSnapshot POSTs to /snapshots and hands back the new record. docs say POST comes back
// empty so if theres no id in the reply poll for it by full name
func (c *Credentials) postSnapshot(config map[string]interface{}, fullName string, timeout int) (*CreateOrUpdateVolumeGroupSnapshotResponse, error) {
	apiRequest, err := c.Post("/snapshots", config, timeout)
	if err != nil {
		return nil, err
	}

	var apiResponse CreateOrUpdateVolumeGroupSnapshotResponse
	if err := mapstructure.Decode(apiRequest, &apiResponse); err == nil && apiResponse.ID != 0 {
		return &apiResponse, nil
	}

	for i := 0; i < 10; i++ {
		snaps, err := c.GetVolumeGroupSnapshotByName(fullName, timeout)
		if err != nil {
			return nil, err
		}
		if len(snaps.Hits) > 0 {
			created := CreateOrUpdateVolumeGroupSnapshotResponse(snaps.Hits[0])
			return &created, nil
		}
		time.Sleep(time.Second)
	}
	return nil, fmt.Errorf("Created '%s' but it never showed up on the server", fullName)
}

// GetVolumeGroupSnapshotByID returns the snapshot or view with the given id (empty Hits if it doesnt exist).
func (c *Credentials) GetVolumeGroupSnapshotByID(id int, timeout ...int) (*GetVolumeGroupSnapshotResponse, error) {
	var apiResponse GetVolumeGroupSnapshotResponse
	if err := c.queryInto("/snapshots", url.Values{"id": {strconv.Itoa(id)}}, &apiResponse, httpTimeout(timeout)); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

// GetVolumeGroupSnapshotByName returns the snapshot or view with the full volumegroup:name.
func (c *Credentials) GetVolumeGroupSnapshotByName(name string, timeout ...int) (*GetVolumeGroupSnapshotResponse, error) {
	var apiResponse GetVolumeGroupSnapshotResponse
	if err := c.queryInto("/snapshots", url.Values{"name": {name}}, &apiResponse, httpTimeout(timeout)); err != nil {
		return nil, err
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

// /snapshots has no PATCH, left here commented in case that changes
// // UpdateVolumeGroupSnapshotByID patches a snapshot or view.
// //
// // Valid config keys are: short_name, retention_policy (pass the policy name, its converted to a ref), is_auto_deleteable.
// func (c *Credentials) UpdateVolumeGroupSnapshotByID(id int, config map[string]interface{}, timeout ...int) (*CreateOrUpdateVolumeGroupSnapshotResponse, error) {
//
// 	httpTimeout := httpTimeout(timeout)
//
// 	validUpdateKeys := []string{"short_name", "retention_policy", "is_auto_deleteable"}
// 	body := map[string]interface{}{}
// 	for key, value := range config {
// 		if c.stringInSlice(validUpdateKeys, key) == false {
// 			return nil, fmt.Errorf("The provided 'config' parameter contains invalid keys. 'short_name', 'retention_policy' and 'is_auto_deleteable' are the only valid choices")
// 		}
// 		body[key] = value
// 	}
//
// 	if rp, ok := config["retention_policy"].(string); ok {
// 		retentionPolicyID, err := c.GetRetentionPolicyID(rp, httpTimeout)
// 		if err != nil {
// 			return nil, err
// 		}
// 		body["retention_policy"] = map[string]string{"ref": ref("retention_policies", retentionPolicyID)}
// 	}
//
// 	if _, err := c.Patch(fmt.Sprintf("/snapshots/%d", id), body, httpTimeout); err != nil {
// 		return nil, err
// 	}
//
// 	snaps, err := c.GetVolumeGroupSnapshotByID(id, httpTimeout)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(snaps.Hits) == 0 {
// 		return nil, fmt.Errorf("Snapshot %d disappeared after update", id)
// 	}
// 	updated := CreateOrUpdateVolumeGroupSnapshotResponse(snaps.Hits[0])
// 	return &updated, nil
// }

// DeleteVolumeGroupSnapshotByID deletes a snapshot or view by id. Unmap views first.
func (c *Credentials) DeleteVolumeGroupSnapshotByID(id int, timeout ...int) (*DeleteResponse, error) {
	apiRequest, err := c.Delete(fmt.Sprintf("/snapshots/%d", id), httpTimeout(timeout))
	if err != nil {
		return nil, err
	}

	var apiResponse DeleteResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}
	return &apiResponse, nil
}
