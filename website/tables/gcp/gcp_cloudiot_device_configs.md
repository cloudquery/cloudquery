# Table: gcp_cloudiot_device_configs

This table shows data for GCP Cloud IoT Device Configs.

https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices.configVersions#DeviceConfig

The composite primary key for this table is (**project_id**, **device_name**, **version**).

## Relations

This table depends on [gcp_cloudiot_devices](gcp_cloudiot_devices).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|device_name (PK)|`utf8`|
|version (PK)|`int64`|
|cloud_update_time|`timestamp[us, tz=UTC]`|
|device_ack_time|`timestamp[us, tz=UTC]`|
|binary_data|`binary`|