# Table: aws_servicediscovery_instances

This table shows data for AWS Cloud Map Instances.

https://docs.aws.amazon.com/cloud-map/latest/api/API_Instance.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **id**).
## Relations

This table depends on [aws_servicediscovery_services](aws_servicediscovery_services.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id|`utf8`|
|attributes|`json`|
|creator_request_id|`utf8`|