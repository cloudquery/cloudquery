# Table: aws_servicediscovery_namespaces

This table shows data for AWS Cloud Map Namespaces.

https://docs.aws.amazon.com/cloud-map/latest/api/API_Namespace.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

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
|id|`utf8`|
|name|`utf8`|
|properties|`json`|
|service_count|`int64`|
|type|`utf8`|