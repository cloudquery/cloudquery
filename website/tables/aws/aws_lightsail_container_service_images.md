# Table: aws_lightsail_container_service_images

This table shows data for Lightsail Container Service Images.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerImage.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_lightsail_container_services](aws_lightsail_container_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|container_service_arn|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|digest|`utf8`|
|image|`utf8`|