# Table: gcp_cloudiot_device_states

https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices.states#DeviceState

The composite primary key for this table is (**project_id**, **device_name**, **update_time**).

## Relations

This table depends on [gcp_cloudiot_devices](gcp_cloudiot_devices.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|device_name (PK)|String|
|update_time (PK)|Timestamp|
|binary_data|ByteArray|