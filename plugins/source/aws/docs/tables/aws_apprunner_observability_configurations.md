# Table: aws_apprunner_observability_configurations

https://docs.aws.amazon.com/apprunner/latest/api/API_ObservabilityConfiguration.html

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
|created_at|Timestamp|
|deleted_at|Timestamp|
|latest|Bool|
|observability_configuration_arn|String|
|observability_configuration_name|String|
|observability_configuration_revision|Int|
|status|String|
|trace_configuration|JSON|