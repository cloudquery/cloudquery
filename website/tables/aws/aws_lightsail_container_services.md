# Table: aws_lightsail_container_services

This table shows data for Lightsail Container Services.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerService.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_lightsail_container_services:
  - [aws_lightsail_container_service_deployments](aws_lightsail_container_service_deployments)
  - [aws_lightsail_container_service_images](aws_lightsail_container_service_images)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|container_service_name|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|current_deployment|`json`|
|is_disabled|`bool`|
|location|`json`|
|next_deployment|`json`|
|power|`utf8`|
|power_id|`utf8`|
|principal_arn|`utf8`|
|private_domain_name|`utf8`|
|private_registry_access|`json`|
|public_domain_names|`json`|
|resource_type|`utf8`|
|scale|`int64`|
|state|`utf8`|
|state_detail|`json`|
|url|`utf8`|