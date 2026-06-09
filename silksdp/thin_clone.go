package silksdp

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// CreateThinClone creates a new thin clone of a Volume on the Silk server. A thin clone is a writable
// volume backed by a Volume Group Snapshot. The resulting object lives at /volumes and can be
// queried, updated, and deleted with the existing Volume functions (GetVolumes, UpdateVolume,
// DeleteVolume).
//
// `volumeGroupName` is the Volume Group that owns the source volume and snapshot.
// `sourceSnapshotName` is the name of the Volume Group Snapshot to clone from.
// `sourceVolumeName` is the name of the Volume within that snapshot to clone.
func (c *Credentials) CreateThinClone(name string, volumeGroupName string, sourceSnapshotName string, sourceVolumeName string, timeout ...int) (*CreateThinCloneResponse, error) {

	httpTimeout := httpTimeout(timeout)

	volumeGroupID, err := c.GetVolumeGroupID(volumeGroupName, httpTimeout)
	if err != nil {
		return nil, err
	}

	snapshotID, err := c.GetVolumeGroupSnapshotID(sourceSnapshotName, httpTimeout)
	if err != nil {
		return nil, err
	}

	sourceVolumeID, err := c.GetVolumeID(sourceVolumeName, httpTimeout)
	if err != nil {
		return nil, err
	}

	volumeGroupRef := map[string]interface{}{"ref": fmt.Sprintf("/volume_groups/%d", volumeGroupID)}
	sourceSnapshotRef := map[string]interface{}{"ref": fmt.Sprintf("/snapshots/%d", snapshotID)}
	sourceVolumeRef := map[string]interface{}{"ref": fmt.Sprintf("/volumes/%d", sourceVolumeID)}

	config := map[string]interface{}{}
	config["name"] = name
	config["volume_group"] = volumeGroupRef
	config["source_snapshot"] = sourceSnapshotRef
	config["source_volume"] = sourceVolumeRef
	config["is_thin_clone"] = "true"

	apiRequest, err := c.Post("/volumes", config, httpTimeout)
	if err != nil {
		return nil, err
	}

	var apiResponse CreateThinCloneResponse
	mapErr := mapstructure.Decode(apiRequest, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	return &apiResponse, nil
}
