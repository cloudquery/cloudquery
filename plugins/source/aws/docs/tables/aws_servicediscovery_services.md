# Table: aws_servicediscovery_services

This table shows data for AWS Cloud Map Services.

https://docs.aws.amazon.com/cloud-map/latest/api/API_Service.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_servicediscovery_services:
  - [aws_servicediscovery_instances](aws_servicediscovery_instances.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|creator_request_id|`utf8`|
|description|`utf8`|
|dns_config|`json`|
|health_check_config|`json`|
|health_check_custom_config|`json`|
|id|`utf8`|
|instance_count|`int64`|
|name|`utf8`|
|namespace_id|`utf8`|
|type|`utf8`|