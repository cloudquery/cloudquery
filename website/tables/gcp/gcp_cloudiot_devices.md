# Table: gcp_cloudiot_devices

This table shows data for GCP Cloud IoT Devices.

https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices#Device

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_cloudiot_device_registries](gcp_cloudiot_device_registries).

The following tables depend on gcp_cloudiot_devices:
  - [gcp_cloudiot_device_configs](gcp_cloudiot_device_configs)
  - [gcp_cloudiot_device_states](gcp_cloudiot_device_states)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|id|`utf8`|
|name (PK)|`utf8`|
|num_id|`int64`|
|credentials|`json`|
|last_heartbeat_time|`timestamp[us, tz=UTC]`|
|last_event_time|`timestamp[us, tz=UTC]`|
|last_state_time|`timestamp[us, tz=UTC]`|
|last_config_ack_time|`timestamp[us, tz=UTC]`|
|last_config_send_time|`timestamp[us, tz=UTC]`|
|blocked|`bool`|
|last_error_time|`timestamp[us, tz=UTC]`|
|last_error_status|`json`|
|config|`json`|
|state|`json`|
|log_level|`utf8`|
|metadata|`json`|
|gateway_config|`json`|