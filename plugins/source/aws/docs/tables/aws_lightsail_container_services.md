# Table: aws_lightsail_container_services

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerService.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_lightsail_container_services:
  - [aws_lightsail_container_service_deployments](aws_lightsail_container_service_deployments.md)
  - [aws_lightsail_container_service_images](aws_lightsail_container_service_images.md)

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
|container_service_name|String|
|created_at|Timestamp|
|current_deployment|JSON|
|is_disabled|Bool|
|location|JSON|
|next_deployment|JSON|
|power|String|
|power_id|String|
|principal_arn|String|
|private_domain_name|String|
|private_registry_access|JSON|
|public_domain_names|JSON|
|resource_type|String|
|scale|Int|
|state|String|
|state_detail|JSON|
|tags|JSON|
|url|String|