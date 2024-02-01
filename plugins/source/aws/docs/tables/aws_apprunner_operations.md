# Table: aws_apprunner_operations

This table shows data for AWS App Runner Operations.

https://docs.aws.amazon.com/apprunner/latest/api/API_OperationSummary.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**service_arn**, **id**).
## Relations

This table depends on [aws_apprunner_services](aws_apprunner_services.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|service_arn|`utf8`|
|ended_at|`timestamp[us, tz=UTC]`|
|id|`utf8`|
|started_at|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|target_arn|`utf8`|
|type|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|