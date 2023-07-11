# Table: aws_servicediscovery_instances

This table shows data for AWS Cloud Map Instances.

https://docs.aws.amazon.com/cloud-map/latest/api/API_Instance.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Relations

This table depends on [aws_servicediscovery_services](aws_servicediscovery_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|id (PK)|`utf8`|
|attributes|`json`|
|creator_request_id|`utf8`|