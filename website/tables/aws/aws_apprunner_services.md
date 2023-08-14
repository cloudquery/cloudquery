# Table: aws_apprunner_services

This table shows data for AWS App Runner Services.

https://docs.aws.amazon.com/apprunner/latest/api/API_Service.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_apprunner_services:
  - [aws_apprunner_custom_domains](aws_apprunner_custom_domains)
  - [aws_apprunner_operations](aws_apprunner_operations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|auto_scaling_configuration_summary|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|instance_configuration|`json`|
|network_configuration|`json`|
|service_arn|`utf8`|
|service_id|`utf8`|
|service_name|`utf8`|
|source_configuration|`json`|
|status|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|deleted_at|`timestamp[us, tz=UTC]`|
|encryption_configuration|`json`|
|health_check_configuration|`json`|
|observability_configuration|`json`|
|service_url|`utf8`|