# Table: aws_apprunner_auto_scaling_configurations

This table shows data for AWS App Runner Auto Scaling Configurations.

https://docs.aws.amazon.com/apprunner/latest/api/API_AutoScalingConfiguration.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|auto_scaling_configuration_arn|`utf8`|
|auto_scaling_configuration_name|`utf8`|
|auto_scaling_configuration_revision|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|deleted_at|`timestamp[us, tz=UTC]`|
|has_associated_service|`bool`|
|is_default|`bool`|
|latest|`bool`|
|max_concurrency|`int64`|
|max_size|`int64`|
|min_size|`int64`|
|status|`utf8`|