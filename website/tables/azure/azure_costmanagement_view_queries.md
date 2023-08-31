# Table: azure_costmanagement_view_queries

This table shows data for Azure Cost Management View Queries.

https://learn.microsoft.com/en-us/rest/api/cost-management/query/usage?tabs=HTTP#queryresult

The primary key for this table is **id**.

## Relations

This table depends on [azure_costmanagement_views](azure_costmanagement_views).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|etag|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|name|`utf8`|
|sku|`utf8`|
|tags|`json`|
|type|`utf8`|