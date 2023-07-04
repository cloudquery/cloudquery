# Table: gcp_baremetalsolution_instances

This table shows data for GCP Bare Metal Solution Instances.

https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.instances#Instance

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|id|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|machine_type|`utf8`|
|state|`utf8`|
|hyperthreading_enabled|`bool`|
|labels|`json`|
|luns|`json`|
|networks|`json`|
|interactive_serial_console_enabled|`bool`|
|os_image|`utf8`|
|pod|`utf8`|
|network_template|`utf8`|
|logical_interfaces|`json`|