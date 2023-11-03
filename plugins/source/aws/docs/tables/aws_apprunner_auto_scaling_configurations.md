# Table: aws_apprunner_auto_scaling_configurations

This table shows data for AWS App Runner Auto Scaling Configurations.

https://docs.aws.amazon.com/apprunner/latest/api/API_AutoScalingConfiguration.html

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
|auto_scaling_configuration_arn|`utf8`|
|auto_scaling_configuration_name|`utf8`|
|auto_scaling_configuration_revision|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|deleted_at|`timestamp[us, tz=UTC]`|
|latest|`bool`|
|max_concurrency|`int64`|
|max_size|`int64`|
|min_size|`int64`|
|status|`utf8`|