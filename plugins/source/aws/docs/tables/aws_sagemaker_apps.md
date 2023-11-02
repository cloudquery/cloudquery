# Table: aws_sagemaker_apps

This table shows data for Amazon SageMaker Apps.

https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeApp.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|app_arn|`utf8`|
|app_name|`utf8`|
|app_type|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|domain_id|`utf8`|
|failure_reason|`utf8`|
|last_health_check_timestamp|`timestamp[us, tz=UTC]`|
|last_user_activity_timestamp|`timestamp[us, tz=UTC]`|
|resource_spec|`json`|
|space_name|`utf8`|
|status|`utf8`|
|user_profile_name|`utf8`|