# Table: gcp_cloudiot_device_states

https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices.states#DeviceState

The primary key for this table is **_cq_id**.

## Relations

This table depends on [gcp_cloudiot_devices](gcp_cloudiot_devices.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|device_name|String|
|update_time|Timestamp|
|binary_data|ByteArray|