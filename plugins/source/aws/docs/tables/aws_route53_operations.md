# Table: aws_route53_operations

This table shows data for Amazon Route 53 Operations.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **operation_id**, **status**, **submitted_date**, **type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|domain_name|`utf8`|
|last_updated_date|`timestamp[us, tz=UTC]`|
|message|`utf8`|
|operation_id|`utf8`|
|status|`utf8`|
|status_flag|`utf8`|
|submitted_date|`timestamp[us, tz=UTC]`|
|type|`utf8`|