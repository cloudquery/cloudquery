# Table: gcp_compute_disks

This table shows data for GCP Compute Disks.

https://cloud.google.com/compute/docs/reference/rest/v1/disks#Disk

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id|utf8|
|architecture|utf8|
|creation_timestamp|utf8|
|description|utf8|
|disk_encryption_key|json|
|guest_os_features|json|
|id|int64|
|kind|utf8|
|label_fingerprint|utf8|
|labels|json|
|last_attach_timestamp|utf8|
|last_detach_timestamp|utf8|
|license_codes|list<item: int64, nullable>|
|licenses|list<item: utf8, nullable>|
|location_hint|utf8|
|name|utf8|
|options|utf8|
|params|json|
|physical_block_size_bytes|int64|
|provisioned_iops|int64|
|region|utf8|
|replica_zones|list<item: utf8, nullable>|
|resource_policies|list<item: utf8, nullable>|
|satisfies_pzs|bool|
|self_link (PK)|utf8|
|size_gb|int64|
|source_disk|utf8|
|source_disk_id|utf8|
|source_image|utf8|
|source_image_encryption_key|json|
|source_image_id|utf8|
|source_snapshot|utf8|
|source_snapshot_encryption_key|json|
|source_snapshot_id|utf8|
|source_storage_object|utf8|
|status|utf8|
|type|utf8|
|users|list<item: utf8, nullable>|
|zone|utf8|