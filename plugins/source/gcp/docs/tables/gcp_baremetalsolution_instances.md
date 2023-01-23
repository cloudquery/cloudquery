# Table: gcp_baremetalsolution_instances

https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.instances#Instance

The composite primary key for this table is (**project_id**, **name**).

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
|create_time|Timestamp|
|update_time|Timestamp|
|machine_type|String|
|state|String|
|hyperthreading_enabled|Bool|
|labels|JSON|
|luns|JSON|
|networks|JSON|
|interactive_serial_console_enabled|Bool|
|os_image|String|
|pod|String|
|network_template|String|
|logical_interfaces|JSON|