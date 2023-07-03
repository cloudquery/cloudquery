# Table: gcp_compute_zones

This table shows data for GCP Compute Zones.

https://cloud.google.com/compute/docs/reference/rest/v1/zones/list#response-body

The primary key for this table is **self_link**.

## Relations

The following tables depend on gcp_compute_zones:
  - [gcp_compute_machine_types](gcp_compute_machine_types)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|available_cpu_platforms|`list<item: utf8, nullable>`|
|creation_timestamp|`utf8`|
|deprecated|`json`|
|description|`utf8`|
|id|`int64`|
|kind|`utf8`|
|name|`utf8`|
|region|`utf8`|
|self_link (PK)|`utf8`|
|status|`utf8`|
|supports_pzs|`bool`|