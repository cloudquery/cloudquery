# Table: gcp_compute_zones

https://cloud.google.com/compute/docs/reference/rest/v1/zones/list#response-body

The primary key for this table is **self_link**.

## Relations

The following tables depend on gcp_compute_zones:
  - [gcp_compute_machine_types](gcp_compute_machine_types.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|available_cpu_platforms|StringArray|
|creation_timestamp|String|
|deprecated|JSON|
|description|String|
|id|Int|
|kind|String|
|name|String|
|region|String|
|self_link (PK)|String|
|status|String|
|supports_pzs|Bool|