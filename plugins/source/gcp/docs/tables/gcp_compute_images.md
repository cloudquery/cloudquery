# Table: gcp_compute_images



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
|archive_size_bytes|Int|
|creation_timestamp|String|
|deprecated|JSON|
|description|String|
|disk_size_gb|Int|
|family|String|
|guest_os_features|JSON|
|id|Int|
|image_encryption_key|JSON|
|kind|String|
|label_fingerprint|String|
|labels|JSON|
|license_codes|IntArray|
|licenses|StringArray|
|name|String|
|raw_disk|JSON|
|satisfies_pzs|Bool|
|shielded_instance_initial_state|JSON|
|source_disk|String|
|source_disk_encryption_key|JSON|
|source_disk_id|String|
|source_image|String|
|source_image_encryption_key|JSON|
|source_image_id|String|
|source_snapshot|String|
|source_snapshot_encryption_key|JSON|
|source_snapshot_id|String|
|source_type|String|
|status|String|
|storage_locations|StringArray|