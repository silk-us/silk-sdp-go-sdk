# v1.4.0

Lookups now let the API do the filtering instead of pulling whole collections and searching them in Go. The same change shipped in the PowerShell SDK in v2.1. Exported function signatures are unchanged, but a few behaviors changed. Read the first section before upgrading.

## Behavior changes

* **Every GET asks for the full set.** Requests now send `__limit=9999`. Before, any list call (`GetVolumes`, `GetHosts`, `GetHostMappings`, ...) silently stopped at the API's default page of 100.
* **`*ByName` functions match exactly.** `GetHostByName`, `GetVolumeByName`, `GetVolumeGroupByName`, `GetHostGroupByName`, `GetRetentionPolicyByName` and `GetCapacityPolicyByName` used `name__contains`, a substring match that the API handles badly. It returned partial matches and could hang on long names. They now send `name=`, which is exact. The API's exact match ignores case. The `Get*ID` functions still require the case to match.
* **`GetVolumeGroupSnapshotID` with a short name** now returns an error when more than one Volume Group has a snapshot with that name. Pass the full `{volume group}:{snapshot}` name to avoid this.
* **`DeleteHost`** returns a `DeleteResponse` with `StatusCode` 404 when the host doesn't exist. Before, it crashed.
* **Empty API responses** now come back as an error, or a status-only response, instead of crashing.
* **The one-second pause after each call** now only happens after POSTs, where the API needs a moment to make the new object visible.

## Snapshots

* `CreateVolumeGroupSnapshot` has been rewritten. The old version sent fields the API doesn't accept and never worked. Its signature is unchanged. Pass `deletable = false` for anything you need to keep, so the Retention Policy can't expire it.
* `CreateOrUpdateVolumeGroupSnapshotResponse` now has the full snapshot record. Its `Source`, `RetentionPolicy` and `VolumeGroup` fields are `{ Ref }` structs instead of strings.
* `GetVolumeGroupSnapshotResponse.Hits` is now a `[]VolumeGroupSnapshot`, a named type. Field access is unchanged.
* `/snapshots` doesn't support PATCH, so there is no update function.

## New functions

* `GetHostByID`, `GetHostGroupByID`, `GetVolumeByID`, `GetVolumeGroupByID`, `GetRetentionPolicyByID`, `GetCapacityPolicyByID`, `GetVolumeGroupSnapshotByID`
* `GetVolumeGroupSnapshotByName`
* `GetHostIQNByID`, `GetHostPWWNByID`, `GetHostGroupHostsByID`: same as the name versions, without the name lookup.
* `GetVolumeMappingsByID`: returns the hosts and host groups mapped to a volume, from one mappings call.
* `GetVolumeGroupSnapshotIDInGroup`: snapshot lookup inside a known Volume Group. Accepts the short or full name.
* `CreateVolumeGroupView`
* `DeleteVolumeGroupSnapshotByID`

## Fixes

* Mapping lookups (`GetVolumeHostMappings`, `GetVolumeHostGroupMappings`, `GetVolumeGroupHostMappings`, `GetVolumeGroupHostGroupMappings`) used to make one extra API call per mapping. They now make one mappings call plus one lookup per object type.
* Query strings are built with `net/url`, so names containing `&`, `+` or spaces are now sent correctly.
* `CreateThinClone` looks the snapshot up inside the clone's Volume Group, so a short snapshot name works.
