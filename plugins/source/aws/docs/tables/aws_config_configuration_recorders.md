# Table: aws_config_configuration_recorders


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|name|String|
|recording_group|JSON|
|role_arn|String|
|status_last_error_code|String|
|status_last_error_message|String|
|status_last_start_time|Timestamp|
|status_last_status|String|
|status_last_status_change_time|Timestamp|
|status_last_stop_time|Timestamp|
|status_recording|Bool|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|