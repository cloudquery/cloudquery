# Table: gcp_compute_images

This table shows data for GCP Compute Images.

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|architecture|`utf8`|
|archive_size_bytes|`int64`|
|creation_timestamp|`utf8`|
|deprecated|`json`|
|description|`utf8`|
|disk_size_gb|`int64`|
|family|`utf8`|
|guest_os_features|`json`|
|id|`int64`|
|image_encryption_key|`json`|
|kind|`utf8`|
|label_fingerprint|`utf8`|
|labels|`json`|
|license_codes|`list<item: int64, nullable>`|
|licenses|`list<item: utf8, nullable>`|
|name|`utf8`|
|raw_disk|`json`|
|satisfies_pzs|`bool`|
|self_link (PK)|`utf8`|
|shielded_instance_initial_state|`json`|
|source_disk|`utf8`|
|source_disk_encryption_key|`json`|
|source_disk_id|`utf8`|
|source_image|`utf8`|
|source_image_encryption_key|`json`|
|source_image_id|`utf8`|
|source_snapshot|`utf8`|
|source_snapshot_encryption_key|`json`|
|source_snapshot_id|`utf8`|
|source_type|`utf8`|
|status|`utf8`|
|storage_locations|`list<item: utf8, nullable>`|