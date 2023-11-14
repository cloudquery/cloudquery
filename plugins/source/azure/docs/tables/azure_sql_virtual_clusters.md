# Table: azure_sql_virtual_clusters

This table shows data for Azure SQL Virtual Clusters.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/virtual-clusters/list?tabs=HTTP#virtualcluster

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|