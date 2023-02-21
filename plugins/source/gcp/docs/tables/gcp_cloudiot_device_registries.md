# Table: gcp_cloudiot_device_registries

https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries#DeviceRegistry

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_cloudiot_device_registries:
  - [gcp_cloudiot_devices](gcp_cloudiot_devices.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|id|String|
|name (PK)|String|
|event_notification_configs|JSON|
|state_notification_config|JSON|
|mqtt_config|JSON|
|http_config|JSON|
|log_level|String|
|credentials|JSON|