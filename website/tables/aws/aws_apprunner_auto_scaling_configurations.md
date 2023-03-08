# Table: aws_apprunner_auto_scaling_configurations

https://docs.aws.amazon.com/apprunner/latest/api/API_AutoScalingConfiguration.html

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
|tags|JSON|
|auto_scaling_configuration_arn|String|
|auto_scaling_configuration_name|String|
|auto_scaling_configuration_revision|Int|
|created_at|Timestamp|
|deleted_at|Timestamp|
|latest|Bool|
|max_concurrency|Int|
|max_size|Int|
|min_size|Int|
|status|String|