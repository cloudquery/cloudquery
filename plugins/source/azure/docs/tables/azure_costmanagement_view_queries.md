# Table: azure_costmanagement_view_queries

https://learn.microsoft.com/en-us/rest/api/cost-management/query/usage?tabs=HTTP#queryresult

The primary key for this table is **id**.

## Relations

This table depends on [azure_costmanagement_views](azure_costmanagement_views.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|etag|String|
|properties|JSON|
|location|String|
|name|String|
|sku|String|
|tags|JSON|
|type|String|