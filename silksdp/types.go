package silksdp

type (
	// GetVolumeGroupsResponse holds the response of the GetVolumeGroups() function
	GetVolumeGroupsResponse struct {
		Hits []struct {
			CapacityPolicy             string      `json:"capacity_policy"`
			CapacityState              string      `json:"capacity_state"`
			CreationTime               float64     `json:"creation_time"`
			Description                interface{} `json:"description"`
			ID                         int         `json:"id"`
			IsDedup                    bool        `json:"is_dedup"`
			IsDefault                  bool        `json:"is_default"`
			IscsiTgtConvertedName      string      `json:"iscsi_tgt_converted_name"`
			LastRestoredFrom           interface{} `json:"last_restored_from"`
			LastRestoredTime           interface{} `json:"last_restored_time"`
			LastSnapshotCreationTime   int         `json:"last_snapshot_creation_time"`
			LogicalCapacity            float64     `json:"logical_capacity"`
			MappedHostsCount           int         `json:"mapped_hosts_count"`
			Name                       string      `json:"name"`
			Quota                      interface{} `json:"quota"`
			ReplicationPeerVolumeGroup interface{} `json:"replication_peer_volume_group"`
			ReplicationRpoHistory      interface{} `json:"replication_rpo_history"`
			ReplicationSession         interface{} `json:"replication_session"`
			SnapshotsCount             int         `json:"snapshots_count"`
			SnapshotsLogicalCapacity   int         `json:"snapshots_logical_capacity"`
			SnapshotsOverheadState     string      `json:"snapshots_overhead_state"`
			ViewsCount                 int         `json:"views_count"`
			VolumesCount               int         `json:"volumes_count"`
			VolumesLogicalCapacity     int         `json:"volumes_logical_capacity"`
			VolumesProvisionedCapacity int64       `json:"volumes_provisioned_capacity"`
		} `json:"hits"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	}

	// CreateOrUpdateVolumeGroupResponse holds the response of the CreateVolumeGroup() and
	// UpdateVolumeGroup() functions
	CreateOrUpdateVolumeGroupResponse struct {
		CapacityPolicy             interface{} `json:"capacity_policy"`
		CapacityState              string      `json:"capacity_state"`
		CreationTime               interface{} `json:"creation_time"`
		Description                interface{} `json:"description"`
		ID                         int         `json:"id"`
		IsDedup                    bool        `json:"is_dedup"`
		IsDefault                  bool        `json:"is_default"`
		LastRestoredFrom           interface{} `json:"last_restored_from"`
		LastRestoredTime           interface{} `json:"last_restored_time"`
		LastSnapshotCreationTime   interface{} `json:"last_snapshot_creation_time"`
		LogicalCapacity            int         `json:"logical_capacity"`
		MappedHostsCount           int         `json:"mapped_hosts_count"`
		Name                       string      `json:"name"`
		Quota                      interface{} `json:"quota"`
		ReplicationPeerVolumeGroup interface{} `json:"replication_peer_volume_group"`
		ReplicationRpoHistory      interface{} `json:"replication_rpo_history"`
		ReplicationSession         interface{} `json:"replication_session"`
		SnapshotsCount             int         `json:"snapshots_count"`
		SnapshotsLogicalCapacity   int         `json:"snapshots_logical_capacity"`
		SnapshotsOverheadState     string      `json:"snapshots_overhead_state"`
		ViewsCount                 int         `json:"views_count"`
		VolumesCount               int         `json:"volumes_count"`
		VolumesLogicalCapacity     int         `json:"volumes_logical_capacity"`
		VolumesProvisionedCapacity int         `json:"volumes_provisioned_capacity"`
		PipeID                     int         `json:"pipeId"`
		PipeName                   string      `json:"pipeName"`
	}

	// CreateOrUpdateVolumeResponse holds the response of the CreateVolume() and
	// UpdateVolume() functions
	CreateOrUpdateVolumeResponse struct {
		AvgCompressedRatio          float64     `json:"avg_compressed_ratio"`
		AvgCompressedRatioTimestamp int         `json:"avg_compressed_ratio_timestamp"`
		CreationTime                int         `json:"creation_time"`
		CurrentReplicationStats     interface{} `json:"current_replication_stats"`
		CurrentStats                struct {
			Ref string `json:"ref"`
		} `json:"current_stats"`
		DedupSource                    int         `json:"dedup_source"`
		DedupTarget                    int         `json:"dedup_target"`
		Description                    interface{} `json:"description"`
		ID                             int         `json:"id"`
		IsDedup                        bool        `json:"is_dedup"`
		IsNew                          bool        `json:"is_new"`
		LastRestoredFrom               interface{} `json:"last_restored_from"`
		LastRestoredTime               interface{} `json:"last_restored_time"`
		LogicalCapacity                int         `json:"logical_capacity"`
		MarkedForDeletion              bool        `json:"marked_for_deletion"`
		Name                           string      `json:"name"`
		NoDedup                        int         `json:"no_dedup"`
		NodeID                         int         `json:"node_id"`
		ReadOnly                       bool        `json:"read_only"`
		ReplicationPeerVolume          interface{} `json:"replication_peer_volume"`
		ScsiSn                         string      `json:"scsi_sn"`
		ScsiSuffix                     int         `json:"scsi_suffix"`
		Size                           int         `json:"size"`
		SnapshotsLogicalCapacity       int         `json:"snapshots_logical_capacity"`
		StreamAvgCompressedSizeInBytes int         `json:"stream_avg_compressed_size_in_bytes"`
		VmwareSupport                  bool        `json:"vmware_support"`
		VolumeGroup                    struct {
			Ref string `json:"ref"`
		} `json:"volume_group"`
		PipeID   int    `json:"pipeId"`
		PipeName string `json:"pipeName"`
	}

	// GetVolumesResponse holds the response of the GetVolumes() function
	GetVolumesResponse struct {
		Hits []struct {
			AvgCompressedRatio          string  `json:"avg_compressed_ratio"`
			AvgCompressedRatioTimestamp float64 `json:"avg_compressed_ratio_timestamp"`
			CreationTime                int     `json:"creation_time"`
			CurrentReplicationStats     string  `json:"current_replication_stats"`
			CurrentStats                struct {
				Ref string `json:"ref"`
			} `json:"current_stats"`
			DedupSource                    int     `json:"dedup_source"`
			DedupTarget                    int     `json:"dedup_target"`
			Description                    string  `json:"description"`
			ID                             int     `json:"id"`
			IsDedup                        bool    `json:"is_dedup"`
			IsNew                          bool    `json:"is_new"`
			LastRestoredFrom               string  `json:"last_restored_from"`
			LastRestoredTime               string  `json:"last_restored_time"`
			LogicalCapacity                float64 `json:"logical_capacity"`
			MarkedForDeletion              bool    `json:"marked_for_deletion"`
			Name                           string  `json:"name"`
			NoDedup                        int     `json:"no_dedup"`
			NodeID                         int     `json:"node_id"`
			ReadOnly                       bool    `json:"read_only"`
			ReplicationPeerVolume          string  `json:"replication_peer_volume"`
			ScsiSn                         string  `json:"scsi_sn"`
			ScsiSuffix                     int     `json:"scsi_suffix"`
			Size                           int     `json:"size"`
			SnapshotsLogicalCapacity       float64 `json:"snapshots_logical_capacity"`
			StreamAvgCompressedSizeInBytes float64 `json:"stream_avg_compressed_size_in_bytes"`
			VmwareSupport                  bool    `json:"vmware_support"`
			VolumeGroup                    struct {
				Ref string `json:"ref"`
			} `json:"volume_group"`
		} `json:"hits"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	}

	// CreateOrUpdateHostResponse holds the response of the CreateHost() and
	// UpdateHost() functions
	CreateOrUpdateHostResponse struct {
		HostGroup     string `json:"host_group"`
		ID            int    `json:"id"`
		IsPartOfGroup bool   `json:"is_part_of_group"`
		Name          string `json:"name"`
		Type          string `json:"type"`
		ViewsCount    int    `json:"views_count"`
		VolumesCount  int    `json:"volumes_count"`
	}

	// GetHostsResponse holds the response of the GetHosts() function
	GetHostsResponse struct {
		Hits []struct {
			HostGroup     interface{} `json:"host_group"`
			ID            int         `json:"id"`
			IsPartOfGroup bool        `json:"is_part_of_group"`
			Name          string      `json:"name"`
			Type          string      `json:"type"`
			ViewsCount    int         `json:"views_count"`
			VolumesCount  int         `json:"volumes_count"`
		} `json:"hits"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	}

	// CreateOrUpdateHostGroupResponse holds the response of the CreateHostGroup() and
	// UpdateHostGroup() functions
	CreateOrUpdateHostGroupResponse struct {
		AllowDifferentHostTypes bool        `json:"allow_different_host_types"`
		Description             interface{} `json:"description"`
		HostsCount              int         `json:"hosts_count"`
		ID                      int         `json:"id"`
		Name                    string      `json:"name"`
		ViewsCount              int         `json:"views_count"`
		VolumesCount            int         `json:"volumes_count"`
	}

	// GetHostGroupsResponse holds the response of the GetHostGroups() function
	GetHostGroupsResponse struct {
		Hits []struct {
			AllowDifferentHostTypes bool        `json:"allow_different_host_types"`
			Description             interface{} `json:"description"`
			HostsCount              int         `json:"hosts_count"`
			ID                      int         `json:"id"`
			Name                    string      `json:"name"`
			ViewsCount              int         `json:"views_count"`
			VolumesCount            int         `json:"volumes_count"`
		} `json:"hits"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	}

	// CreateHostVolumeMappingResponse holds the response of the CreateHostVolumeMapping() function
	CreateHostVolumeMappingResponse struct {
		Host struct {
			Ref string `json:"ref"`
		} `json:"host"`
		ID     int `json:"id"`
		Lun    int `json:"lun"`
		Volume struct {
			Ref string `json:"ref"`
		} `json:"volume"`
	}

	// GetHostMappingsResponse holds the response of the GET /mappings API call used inside the GetHostMappings()
	// function. This value is then filtered to returned the "Hits" responses in IndividualHostMappingResponse
	GetHostMappingsResponse struct {
		Hits []struct {
			Host struct {
				Ref string `json:"ref"`
			} `json:"host"`
			ID     int `json:"id"`
			Lun    int `json:"lun,omitempty"`
			Volume struct {
				Ref string `json:"ref"`
			} `json:"volume"`
		} `json:"hits"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	}

	// IndividualHostMappingResponse holds Host mappings returned by the GetHostMappingsResponse() function
	IndividualHostMappingResponse struct {
		Host struct {
			Ref string "json:\"ref\""
		} "json:\"host\""
		ID     int "json:\"id\""
		Lun    int "json:\"lun,omitempty\""
		Volume struct {
			Ref string "json:\"ref\""
		} "json:\"volume\""
	}

	// GetHostPWWNResponse holds the response of the GET /host_fc_ports API call used inside the GetHostPWWN()
	// function. This value is then filtered to returned the "Hits" responses in IndividualHostPWWNResponse
	GetHostPWWNResponse struct {
		Hits []struct {
			Host struct {
				Ref string `json:"ref"`
			} `json:"host"`
			ID   int    `json:"id"`
			Pwwn string `json:"pwwn"`
		} `json:"hits"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	}

	// IndividualHostPWWNResponse holds the PWWN and Host mappings returned by the GetHostPWWN() function
	IndividualHostPWWNResponse struct {
		Host struct {
			Ref string `json:"ref"`
		} `json:"host"`
		ID   int    `json:"id"`
		Pwwn string `json:"pwwn"`
	}

	// CreateHostPWWNResponse holds the response of the CreateHostPWWN() function
	CreateHostPWWNResponse struct {
		Host struct {
			Ref string `json:"ref"`
		} `json:"host"`
		ID   int    `json:"id"`
		Pwwn string `json:"pwwn"`
	}

	// GetHostIQNResponse holds the response of the GET /host_iqns API call used inside the GetHostIQN()
	// function. This value is then filtered to returned the "Hits" responses in IndividualHostIQNResponse
	GetHostIQNResponse struct {
		Hits []struct {
			Host struct {
				Ref string `json:"ref"`
			} `json:"host"`
			ID  int    `json:"id"`
			Iqn string `json:"iqn"`
		} `json:"hits"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	}

	// IndividualHostIQNResponse holds the IQN and Host mappings returned by the GetHostIQN() function
	IndividualHostIQNResponse struct {
		Host struct {
			Ref string `json:"ref"`
		} `json:"host"`
		ID  int    `json:"id"`
		Iqn string `json:"iqn"`
	}

	// CreateHostIQNResponse holds the response of the CreateHostIQN() function
	CreateHostIQNResponse struct {
		Host struct {
			Ref string `json:"ref"`
		} `json:"host"`
		ID  int    `json:"id"`
		Iqn string `json:"iqn"`
	}

	// DeleteResponse holds the response of the Delete base function. The status code will always be 204.
	DeleteResponse struct {
		StatusCode int `json:"status_code"`
	}
)
