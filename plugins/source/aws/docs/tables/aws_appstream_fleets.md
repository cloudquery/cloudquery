# Table: aws_appstream_fleets

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Fleet.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|compute_capacity_status|JSON|
|instance_type|String|
|name|String|
|state|String|
|created_time|Timestamp|
|description|String|
|disconnect_timeout_in_seconds|Int|
|display_name|String|
|domain_join_info|JSON|
|enable_default_internet_access|Bool|
|fleet_errors|JSON|
|fleet_type|String|
|iam_role_arn|String|
|idle_disconnect_timeout_in_seconds|Int|
|image_arn|String|
|image_name|String|
|max_concurrent_sessions|Int|
|max_user_duration_in_seconds|Int|
|platform|String|
|session_script_s3_location|JSON|
|stream_view|String|
|usb_device_filter_strings|StringArray|
|vpc_config|JSON|