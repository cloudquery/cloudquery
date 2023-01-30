# Table: azure_costmanagement_views

https://learn.microsoft.com/en-us/rest/api/cost-management/views/list?tabs=HTTP#view

The primary key for this table is **id**.

## Relations

The following tables depend on azure_costmanagement_views:
  - [azure_costmanagement_view_queries](azure_costmanagement_view_queries.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|etag|String|
|properties|JSON|
|id (PK)|String|
|name|String|
|type|String|