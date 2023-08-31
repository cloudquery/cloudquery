# Table: gcp_baremetalsolution_volumes

This table shows data for GCP Bare Metal Solution Volumes.

https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.volumes#Volume

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_baremetalsolution_volumes:
  - [gcp_baremetalsolution_volume_luns](gcp_baremetalsolution_volume_luns)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|id|`utf8`|
|storage_type|`utf8`|
|state|`utf8`|
|requested_size_gib|`int64`|
|current_size_gib|`int64`|
|emergency_size_gib|`int64`|
|auto_grown_size_gib|`int64`|
|remaining_space_gib|`int64`|
|snapshot_reservation_detail|`json`|
|snapshot_auto_delete_behavior|`utf8`|
|labels|`json`|
|snapshot_enabled|`bool`|
|pod|`utf8`|