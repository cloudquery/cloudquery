# Table: azure_costmanagement_views

This table shows data for Azure Cost Management Views.

https://learn.microsoft.com/en-us/rest/api/cost-management/views/list?tabs=HTTP#view

The primary key for this table is **id**.

## Relations

The following tables depend on azure_costmanagement_views:
  - [azure_costmanagement_view_queries](azure_costmanagement_view_queries)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|etag|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|