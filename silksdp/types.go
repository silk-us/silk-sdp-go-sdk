package silksdp

type (
	// GetVolumeGroupsResponse holds the response of the GetVolumeGroups() function
	GetVolumeGroupsResponse struct {
		Hits []struct {
			CapacityPolicy             string      `mapstructure:"capacity_policy"`
			CapacityState              string      `mapstructure:"capacity_state"`
			CreationTime               float64     `mapstructure:"creation_time"`
			Description                interface{} `mapstructure:"description"`
			ID                         int         `mapstructure:"id"`
			IsDedup                    bool        `mapstructure:"is_dedup"`
			IsDefault                  bool        `mapstructure:"is_default"`
			IscsiTgtConvertedName      string      `mapstructure:"iscsi_tgt_converted_name"`
			LastRestoredFrom           interface{} `mapstructure:"last_restored_from"`
			LastRestoredTime           interface{} `mapstructure:"last_restored_time"`
			LastSnapshotCreationTime   int         `mapstructure:"last_snapshot_creation_time"`
			LogicalCapacity            float64     `mapstructure:"logical_capacity"`
			MappedHostsCount           int         `mapstructure:"mapped_hosts_count"`
			Name                       string      `mapstructure:"name"`
			Quota                      interface{} `mapstructure:"quota"`
			ReplicationPeerVolumeGroup interface{} `mapstructure:"replication_peer_volume_group"`
			ReplicationRpoHistory      interface{} `mapstructure:"replication_rpo_history"`
			ReplicationSession         interface{} `mapstructure:"replication_session"`
			SnapshotsCount             int         `mapstructure:"snapshots_count"`
			SnapshotsLogicalCapacity   int         `mapstructure:"snapshots_logical_capacity"`
			SnapshotsOverheadState     string      `mapstructure:"snapshots_overhead_state"`
			ViewsCount                 int         `mapstructure:"views_count"`
			VolumesCount               int         `mapstructure:"volumes_count"`
			VolumesLogicalCapacity     int         `mapstructure:"volumes_logical_capacity"`
			VolumesProvisionedCapacity int64       `mapstructure:"volumes_provisioned_capacity"`
		} `mapstructure:"hits"`
		Limit  int `mapstructure:"limit"`
		Offset int `mapstructure:"offset"`
		Total  int `mapstructure:"total"`
	}

	// CreateOrUpdateVolumeGroupResponse holds the response of the CreateVolumeGroup() and
	// UpdateVolumeGroup() functions
	CreateOrUpdateVolumeGroupResponse struct {
		CapacityPolicy             interface{} `mapstructure:"capacity_policy"`
		CapacityState              string      `mapstructure:"capacity_state"`
		CreationTime               interface{} `mapstructure:"creation_time"`
		Description                interface{} `mapstructure:"description"`
		ID                         int         `mapstructure:"id"`
		IsDedup                    bool        `mapstructure:"is_dedup"`
		IsDefault                  bool        `mapstructure:"is_default"`
		LastRestoredFrom           interface{} `mapstructure:"last_restored_from"`
		LastRestoredTime           interface{} `mapstructure:"last_restored_time"`
		LastSnapshotCreationTime   interface{} `mapstructure:"last_snapshot_creation_time"`
		LogicalCapacity            int         `mapstructure:"logical_capacity"`
		MappedHostsCount           int         `mapstructure:"mapped_hosts_count"`
		Name                       string      `mapstructure:"name"`
		Quota                      interface{} `mapstructure:"quota"`
		ReplicationPeerVolumeGroup interface{} `mapstructure:"replication_peer_volume_group"`
		ReplicationRpoHistory      interface{} `mapstructure:"replication_rpo_history"`
		ReplicationSession         interface{} `mapstructure:"replication_session"`
		SnapshotsCount             int         `mapstructure:"snapshots_count"`
		SnapshotsLogicalCapacity   int         `mapstructure:"snapshots_logical_capacity"`
		SnapshotsOverheadState     string      `mapstructure:"snapshots_overhead_state"`
		ViewsCount                 int         `mapstructure:"views_count"`
		VolumesCount               int         `mapstructure:"volumes_count"`
		VolumesLogicalCapacity     int         `mapstructure:"volumes_logical_capacity"`
		VolumesProvisionedCapacity int         `mapstructure:"volumes_provisioned_capacity"`
		PipeID                     int         `mapstructure:"pipeId"`
		PipeName                   string      `mapstructure:"pipeName"`
	}

	// CreateOrUpdateVolumeResponse holds the response of the CreateVolume() and
	// UpdateVolume() functions
	CreateOrUpdateVolumeResponse struct {
		AvgCompressedRatio          float64     `mapstructure:"avg_compressed_ratio"`
		AvgCompressedRatioTimestamp int         `mapstructure:"avg_compressed_ratio_timestamp"`
		CreationTime                int         `mapstructure:"creation_time"`
		CurrentReplicationStats     interface{} `mapstructure:"current_replication_stats"`
		CurrentStats                struct {
			Ref string `mapstructure:"ref"`
		} `mapstructure:"current_stats"`
		DedupSource                    int         `mapstructure:"dedup_source"`
		DedupTarget                    int         `mapstructure:"dedup_target"`
		Description                    interface{} `mapstructure:"description"`
		ID                             int         `mapstructure:"id"`
		IsDedup                        bool        `mapstructure:"is_dedup"`
		IsNew                          bool        `mapstructure:"is_new"`
		LastRestoredFrom               interface{} `mapstructure:"last_restored_from"`
		LastRestoredTime               interface{} `mapstructure:"last_restored_time"`
		LogicalCapacity                int         `mapstructure:"logical_capacity"`
		MarkedForDeletion              bool        `mapstructure:"marked_for_deletion"`
		Name                           string      `mapstructure:"name"`
		NoDedup                        int         `mapstructure:"no_dedup"`
		NodeID                         int         `mapstructure:"node_id"`
		ReadOnly                       bool        `mapstructure:"read_only"`
		ReplicationPeerVolume          interface{} `mapstructure:"replication_peer_volume"`
		ScsiSn                         string      `mapstructure:"scsi_sn"`
		ScsiSuffix                     int         `mapstructure:"scsi_suffix"`
		Size                           int         `mapstructure:"size"`
		SnapshotsLogicalCapacity       int         `mapstructure:"snapshots_logical_capacity"`
		StreamAvgCompressedSizeInBytes int         `mapstructure:"stream_avg_compressed_size_in_bytes"`
		VmwareSupport                  bool        `mapstructure:"vmware_support"`
		VolumeGroup                    struct {
			Ref string `mapstructure:"ref"`
		} `mapstructure:"volume_group"`
		PipeID   int    `mapstructure:"pipeId"`
		PipeName string `mapstructure:"pipeName"`
	}

	// GetVolumesResponse holds the response of the GetVolumes() function
	GetVolumesResponse struct {
		Hits []struct {
			AvgCompressedRatio          int     `mapstructure:"avg_compressed_ratio"`
			AvgCompressedRatioTimestamp float64 `mapstructure:"avg_compressed_ratio_timestamp"`
			CreationTime                int     `mapstructure:"creation_time"`
			CurrentReplicationStats     struct {
				Ref string `mapstructure:"ref"`
			} `mapstructure:"current_replication_stats"`
			CurrentStats struct {
				Ref string `mapstructure:"ref"`
			} `mapstructure:"current_stats"`
			DedupSource           int    `mapstructure:"dedup_source"`
			DedupTarget           int    `mapstructure:"dedup_target"`
			Description           string `mapstructure:"description"`
			ID                    int    `mapstructure:"id"`
			IsDedup               bool   `mapstructure:"is_dedup"`
			IsNew                 bool   `mapstructure:"is_new"`
			LastRestoredFrom      int    `mapstructure:"last_restored_from"`
			LastRestoredTime      int    `mapstructure:"last_restored_time"`
			LogicalCapacity       int    `mapstructure:"logical_capacity"`
			MarkedForDeletion     bool   `mapstructure:"marked_for_deletion"`
			Name                  string `mapstructure:"name"`
			NoDedup               int    `mapstructure:"no_dedup"`
			NodeID                int    `mapstructure:"node_id"`
			ReadOnly              bool   `mapstructure:"read_only"`
			ReplicationPeerVolume struct {
				Ref string `mapstructure:"ref"`
			} `mapstructure:"replication_peer_volume"`
			ScsiSn                         string `mapstructure:"scsi_sn"`
			ScsiSuffix                     int    `mapstructure:"scsi_suffix"`
			Size                           int    `mapstructure:"size"`
			SnapshotsLogicalCapacity       int    `mapstructure:"snapshots_logical_capacity"`
			StreamAvgCompressedSizeInBytes int    `mapstructure:"stream_avg_compressed_size_in_bytes"`
			VmwareSupport                  bool   `mapstructure:"vmware_support"`
			VolumeGroup                    struct {
				Ref string `mapstructure:"ref"`
			} `mapstructure:"volume_group"`
		} `mapstructure:"hits"`
		Limit  int `mapstructure:"limit"`
		Offset int `mapstructure:"offset"`
		Total  int `mapstructure:"total"`
	}

	// CreateOrUpdateHostResponse holds the response of the CreateHost() and
	// UpdateHost() functions
	CreateOrUpdateHostResponse struct {
		HostGroup     string `mapstructure:"host_group"`
		ID            int    `mapstructure:"id"`
		IsPartOfGroup bool   `mapstructure:"is_part_of_group"`
		Name          string `mapstructure:"name"`
		Type          string `mapstructure:"type"`
		ViewsCount    int    `mapstructure:"views_count"`
		VolumesCount  int    `mapstructure:"volumes_count"`
	}

	// GetHostsResponse holds the response of the GetHosts() function
	GetHostsResponse struct {
		Hits []struct {
			HostGroup     interface{} `mapstructure:"host_group"`
			ID            int         `mapstructure:"id"`
			IsPartOfGroup bool        `mapstructure:"is_part_of_group"`
			Name          string      `mapstructure:"name"`
			Type          string      `mapstructure:"type"`
			ViewsCount    int         `mapstructure:"views_count"`
			VolumesCount  int         `mapstructure:"volumes_count"`
		} `mapstructure:"hits"`
		Limit  int `mapstructure:"limit"`
		Offset int `mapstructure:"offset"`
		Total  int `mapstructure:"total"`
	}

	// CreateOrUpdateHostGroupResponse holds the response of the CreateHostGroup() and
	// UpdateHostGroup() functions
	CreateOrUpdateHostGroupResponse struct {
		AllowDifferentHostTypes bool        `mapstructure:"allow_different_host_types"`
		Description             interface{} `mapstructure:"description"`
		HostsCount              int         `mapstructure:"hosts_count"`
		ID                      int         `mapstructure:"id"`
		Name                    string      `mapstructure:"name"`
		ViewsCount              int         `mapstructure:"views_count"`
		VolumesCount            int         `mapstructure:"volumes_count"`
	}

	// GetHostGroupsResponse holds the response of the GetHostGroups() function
	GetHostGroupsResponse struct {
		Hits []struct {
			AllowDifferentHostTypes bool        `mapstructure:"allow_different_host_types"`
			Description             interface{} `mapstructure:"description"`
			HostsCount              int         `mapstructure:"hosts_count"`
			ID                      int         `mapstructure:"id"`
			Name                    string      `mapstructure:"name"`
			ViewsCount              int         `mapstructure:"views_count"`
			VolumesCount            int         `mapstructure:"volumes_count"`
		} `mapstructure:"hits"`
		Limit  int `mapstructure:"limit"`
		Offset int `mapstructure:"offset"`
		Total  int `mapstructure:"total"`
	}

	// CreateHostVolumeMappingResponse holds the response of the CreateHostVolumeMapping() function
	CreateHostVolumeMappingResponse struct {
		Host struct {
			Ref string `mapstructure:"ref"`
		} `mapstructure:"host"`
		ID     int `mapstructure:"id"`
		Lun    int `mapstructure:"lun"`
		Volume struct {
			Ref string `mapstructure:"ref"`
		} `mapstructure:"volume"`
	}

	// GetHostMappingsResponse holds the response of the GET /mappings API call used inside the GetHostMappings()
	// function. This value is then filtered to returned the "Hits" responses in IndividualHostMappingResponse
	GetHostMappingsResponse struct {
		Hits []struct {
			Host struct {
				Ref string `mapstructure:"ref"`
			} `mapstructure:"host"`
			ID     int `mapstructure:"id"`
			Lun    int `mapstructure:"lun,omitempty"`
			Volume struct {
				Ref string `mapstructure:"ref"`
			} `mapstructure:"volume"`
		} `mapstructure:"hits"`
		Limit  int `mapstructure:"limit"`
		Offset int `mapstructure:"offset"`
		Total  int `mapstructure:"total"`
	}

	// IndividualHostMappingResponse holds Host mappings returned by the GetHostMappingsResponse() function
	IndividualHostMappingResponse struct {
		Host struct {
			Ref string "mapstructure:\"ref\""
		} "mapstructure:\"host\""
		ID     int "mapstructure:\"id\""
		Lun    int "mapstructure:\"lun,omitempty\""
		Volume struct {
			Ref string "mapstructure:\"ref\""
		} "mapstructure:\"volume\""
	}

	// GetHostPWWNResponse holds the response of the GET /host_fc_ports API call used inside the GetHostPWWN()
	// function. This value is then filtered to returned the "Hits" responses in IndividualHostPWWNResponse
	GetHostPWWNResponse struct {
		Hits []struct {
			Host struct {
				Ref string `mapstructure:"ref"`
			} `mapstructure:"host"`
			ID   int    `mapstructure:"id"`
			Pwwn string `mapstructure:"pwwn"`
		} `mapstructure:"hits"`
		Limit  int `mapstructure:"limit"`
		Offset int `mapstructure:"offset"`
		Total  int `mapstructure:"total"`
	}

	// IndividualHostPWWNResponse holds the PWWN and Host mappings returned by the GetHostPWWN() function
	IndividualHostPWWNResponse struct {
		Host struct {
			Ref string `mapstructure:"ref"`
		} `mapstructure:"host"`
		ID   int    `mapstructure:"id"`
		Pwwn string `mapstructure:"pwwn"`
	}

	// CreateHostPWWNResponse holds the response of the CreateHostPWWN() function
	CreateHostPWWNResponse struct {
		Host struct {
			Ref string `mapstructure:"ref"`
		} `mapstructure:"host"`
		ID   int    `mapstructure:"id"`
		Pwwn string `mapstructure:"pwwn"`
	}

	// GetHostIQNResponse holds the response of the GET /host_iqns API call used inside the GetHostIQN()
	// function. This value is then filtered to returned the "Hits" responses in IndividualHostIQNResponse
	GetHostIQNResponse struct {
		Hits []struct {
			Host struct {
				Ref string `mapstructure:"ref"`
			} `mapstructure:"host"`
			ID  int    `mapstructure:"id"`
			Iqn string `mapstructure:"iqn"`
		} `mapstructure:"hits"`
		Limit  int `mapstructure:"limit"`
		Offset int `mapstructure:"offset"`
		Total  int `mapstructure:"total"`
	}

	// IndividualHostIQNResponse holds the IQN and Host mappings returned by the GetHostIQN() function
	IndividualHostIQNResponse struct {
		Host struct {
			Ref string `mapstructure:"ref"`
		} `mapstructure:"host"`
		ID  int    `mapstructure:"id"`
		Iqn string `mapstructure:"iqn"`
	}

	// CreateHostIQNResponse holds the response of the CreateHostIQN() function
	CreateHostIQNResponse struct {
		Host struct {
			Ref string `mapstructure:"ref"`
		} `mapstructure:"host"`
		ID  int    `mapstructure:"id"`
		Iqn string `mapstructure:"iqn"`
	}

	// DeleteResponse holds the response of the Delete base function. The status code will always be 204.
	DeleteResponse struct {
		StatusCode int `mapstructure:"status_code"`
	}
)
