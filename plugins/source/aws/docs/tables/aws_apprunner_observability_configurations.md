# Table: aws_apprunner_observability_configurations

This table shows data for AWS App Runner Observability Configurations.

https://docs.aws.amazon.com/apprunner/latest/api/API_ObservabilityConfiguration.html

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
|created_at|`timestamp[us, tz=UTC]`|
|deleted_at|`timestamp[us, tz=UTC]`|
|latest|`bool`|
|observability_configuration_arn|`utf8`|
|observability_configuration_name|`utf8`|
|observability_configuration_revision|`int64`|
|status|`utf8`|
|trace_configuration|`json`|