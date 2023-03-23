# Table: azure_compute_skus

This table shows data for Azure Compute Skus.

https://learn.microsoft.com/en-us/rest/api/compute/resource-skus/list?tabs=HTTP#resourceskusresult

The composite primary key for this table is (**subscription_id**, **locations_hash**, **capabilities_hash**, **name**, **resource_type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id (PK)|String|
|locations_hash (PK)|String|
|capabilities_hash (PK)|String|
|api_versions|StringArray|
|capabilities|JSON|
|capacity|JSON|
|costs|JSON|
|family|String|
|kind|String|
|location_info|JSON|
|locations|StringArray|
|name (PK)|String|
|resource_type (PK)|String|
|restrictions|JSON|
|size|String|
|tier|String|