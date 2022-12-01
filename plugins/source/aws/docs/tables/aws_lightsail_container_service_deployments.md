# Table: aws_lightsail_container_service_deployments

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerServiceDeployment.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_lightsail_container_services](aws_lightsail_container_services.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|container_service_arn|String|
|containers|JSON|
|created_at|Timestamp|
|public_endpoint|JSON|
|state|String|
|version|Int|