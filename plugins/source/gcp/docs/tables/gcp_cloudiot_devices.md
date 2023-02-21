# Table: gcp_cloudiot_devices

https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices#Device

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_cloudiot_device_registries](gcp_cloudiot_device_registries.md).

The following tables depend on gcp_cloudiot_devices:
  - [gcp_cloudiot_device_configs](gcp_cloudiot_device_configs.md)
  - [gcp_cloudiot_device_states](gcp_cloudiot_device_states.md)

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
|num_id|Int|
|credentials|JSON|
|last_heartbeat_time|Timestamp|
|last_event_time|Timestamp|
|last_state_time|Timestamp|
|last_config_ack_time|Timestamp|
|last_config_send_time|Timestamp|
|blocked|Bool|
|last_error_time|Timestamp|
|last_error_status|JSON|
|config|JSON|
|state|JSON|
|log_level|String|
|metadata|JSON|
|gateway_config|JSON|