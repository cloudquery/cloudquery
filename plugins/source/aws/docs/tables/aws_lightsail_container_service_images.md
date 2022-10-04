# Table: aws_lightsail_container_service_images



The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_lightsail_container_services`](aws_lightsail_container_services.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|container_service_arn|String|
|created_at|Timestamp|
|digest|String|
|image|String|