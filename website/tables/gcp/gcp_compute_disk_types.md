# Table: gcp_compute_disk_types

This table shows data for GCP Compute Disk Types.

https://cloud.google.com/compute/docs/reference/rest/v1/diskTypes#DiskType

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|creation_timestamp|`utf8`|
|default_disk_size_gb|`int64`|
|deprecated|`json`|
|description|`utf8`|
|id|`int64`|
|kind|`utf8`|
|name|`utf8`|
|region|`utf8`|
|self_link (PK)|`utf8`|
|valid_disk_size|`utf8`|
|zone|`utf8`|