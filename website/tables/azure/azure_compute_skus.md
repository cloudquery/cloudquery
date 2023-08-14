# Table: azure_compute_skus

This table shows data for Azure Compute Skus.

https://learn.microsoft.com/en-us/rest/api/compute/resource-skus/list?tabs=HTTP#resourceskusresult

The composite primary key for this table is (**subscription_id**, **_sku_hash**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id (PK)|`utf8`|
|_sku_hash (PK)|`utf8`|
|api_versions|`list<item: utf8, nullable>`|
|capabilities|`json`|
|capacity|`json`|
|costs|`json`|
|family|`utf8`|
|kind|`utf8`|
|location_info|`json`|
|locations|`list<item: utf8, nullable>`|
|name (PK)|`utf8`|
|resource_type|`utf8`|
|restrictions|`json`|
|size|`utf8`|
|tier|`utf8`|