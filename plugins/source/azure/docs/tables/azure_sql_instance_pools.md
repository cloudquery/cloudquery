# Table: azure_sql_instance_pools

This table shows data for Azure SQL Instance Pools.

https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/instance-pools/list?tabs=HTTP#instancepool

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|