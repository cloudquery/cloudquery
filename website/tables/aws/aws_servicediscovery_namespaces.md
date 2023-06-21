# Table: aws_servicediscovery_namespaces

This table shows data for Servicediscovery Namespaces.

https://docs.aws.amazon.com/cloud-map/latest/api/API_ListInstances.html

The primary key for this table is **arn**.

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
|create_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|id|`utf8`|
|name|`utf8`|
|properties|`json`|
|service_count|`int64`|
|type|`utf8`|