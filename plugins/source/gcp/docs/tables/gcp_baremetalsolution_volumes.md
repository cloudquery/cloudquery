# Table: gcp_baremetalsolution_volumes

https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.volumes#Volume

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_baremetalsolution_volumes:
  - [gcp_baremetalsolution_volume_luns](gcp_baremetalsolution_volume_luns.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|id|String|
|storage_type|String|
|state|String|
|requested_size_gib|Int|
|current_size_gib|Int|
|emergency_size_gib|Int|
|auto_grown_size_gib|Int|
|remaining_space_gib|Int|
|snapshot_reservation_detail|JSON|
|snapshot_auto_delete_behavior|String|
|labels|JSON|
|snapshot_enabled|Bool|
|pod|String|