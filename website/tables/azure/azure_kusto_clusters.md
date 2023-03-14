# Table: azure_kusto_clusters

This table shows data for Azure Kusto Clusters.

https://learn.microsoft.com/en-us/rest/api/azurerekusto/clusters/list?tabs=HTTP#cluster

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|sku|JSON|
|identity|JSON|
|properties|JSON|
|tags|JSON|
|zones|StringArray|
|etag|String|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|