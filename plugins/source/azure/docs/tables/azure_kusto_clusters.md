# Table: azure_kusto_clusters

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
|id (PK)|String|
|location|String|
|sku|JSON|
|identity|JSON|
|properties|JSON|
|tags|JSON|
|zones|StringArray|
|etag|String|
|name|String|
|system_data|JSON|
|type|String|