# Table: gcp_compute_disks



The primary key for this table is **self_link**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|self_link (PK)|String|
|architecture|String|
|creation_timestamp|String|
|description|String|
|disk_encryption_key|JSON|
|guest_os_features|JSON|
|id|Int|
|kind|String|
|label_fingerprint|String|
|labels|JSON|
|last_attach_timestamp|String|
|last_detach_timestamp|String|
|license_codes|IntArray|
|licenses|StringArray|
|location_hint|String|
|name|String|
|options|String|
|params|JSON|
|physical_block_size_bytes|Int|
|provisioned_iops|Int|
|region|String|
|replica_zones|StringArray|
|resource_policies|StringArray|
|satisfies_pzs|Bool|
|size_gb|Int|
|source_disk|String|
|source_disk_id|String|
|source_image|String|
|source_image_encryption_key|JSON|
|source_image_id|String|
|source_snapshot|String|
|source_snapshot_encryption_key|JSON|
|source_snapshot_id|String|
|source_storage_object|String|
|status|String|
|type|String|
|users|StringArray|
|zone|String|