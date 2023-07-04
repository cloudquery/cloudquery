# Table: gcp_cloudiot_device_registries

This table shows data for GCP Cloud IoT Device Registries.

https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries#DeviceRegistry

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_cloudiot_device_registries:
  - [gcp_cloudiot_devices](gcp_cloudiot_devices)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|id|`utf8`|
|name (PK)|`utf8`|
|event_notification_configs|`json`|
|state_notification_config|`json`|
|mqtt_config|`json`|
|http_config|`json`|
|log_level|`utf8`|
|credentials|`json`|