# Table: aws_route53_operations

This table shows data for Amazon Route 53 Operations.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html

The composite primary key for this table is (**account_id**, **operation_id**, **status**, **submitted_date**, **type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|domain_name|`utf8`|
|last_updated_date|`timestamp[us, tz=UTC]`|
|message|`utf8`|
|operation_id (PK)|`utf8`|
|status (PK)|`utf8`|
|status_flag|`utf8`|
|submitted_date (PK)|`timestamp[us, tz=UTC]`|
|type (PK)|`utf8`|