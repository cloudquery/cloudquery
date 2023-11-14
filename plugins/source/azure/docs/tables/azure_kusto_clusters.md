# Table: azure_kusto_clusters

This table shows data for Azure Kusto Clusters.

https://learn.microsoft.com/en-us/rest/api/azurerekusto/clusters/list?tabs=HTTP#cluster

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|sku|`json`|
|identity|`json`|
|properties|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|etag|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|