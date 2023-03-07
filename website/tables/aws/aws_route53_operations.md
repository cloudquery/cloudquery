# Table: aws_route53_operations

https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html

The composite primary key for this table is (**account_id**, **operation_id**, **status**, **submitted_date**, **type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|domain_name|String|
|last_updated_date|Timestamp|
|message|String|
|operation_id (PK)|String|
|status (PK)|String|
|status_flag|String|
|submitted_date (PK)|Timestamp|
|type (PK)|String|