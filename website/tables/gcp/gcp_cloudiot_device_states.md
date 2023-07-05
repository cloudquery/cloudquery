# Table: gcp_cloudiot_device_states

This table shows data for GCP Cloud IoT Device States.

https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices.states#DeviceState

The composite primary key for this table is (**project_id**, **device_name**, **update_time**).

## Relations

This table depends on [gcp_cloudiot_devices](gcp_cloudiot_devices).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|device_name (PK)|`utf8`|
|update_time (PK)|`timestamp[us, tz=UTC]`|
|binary_data|`binary`|