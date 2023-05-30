# Table: aws_config_configuration_recorders

This table shows data for Config Configuration Recorders.

https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigurationRecorder.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|name|`utf8`|
|recording_group|`json`|
|role_arn|`utf8`|
|status_last_error_code|`utf8`|
|status_last_error_message|`utf8`|
|status_last_start_time|`timestamp[us, tz=UTC]`|
|status_last_status|`utf8`|
|status_last_status_change_time|`timestamp[us, tz=UTC]`|
|status_last_stop_time|`timestamp[us, tz=UTC]`|
|status_recording|`bool`|