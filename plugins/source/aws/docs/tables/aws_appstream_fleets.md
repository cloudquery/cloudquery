# Table: aws_appstream_fleets

This table shows data for Amazon AppStream Fleets.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Fleet.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|compute_capacity_status|`json`|
|instance_type|`utf8`|
|name|`utf8`|
|state|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|disconnect_timeout_in_seconds|`int64`|
|display_name|`utf8`|
|domain_join_info|`json`|
|enable_default_internet_access|`bool`|
|fleet_errors|`json`|
|fleet_type|`utf8`|
|iam_role_arn|`utf8`|
|idle_disconnect_timeout_in_seconds|`int64`|
|image_arn|`utf8`|
|image_name|`utf8`|
|max_concurrent_sessions|`int64`|
|max_user_duration_in_seconds|`int64`|
|platform|`utf8`|
|session_script_s3_location|`json`|
|stream_view|`utf8`|
|usb_device_filter_strings|`list<item: utf8, nullable>`|
|vpc_config|`json`|