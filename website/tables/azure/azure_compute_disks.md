# Table: azure_compute_disks

This table shows data for Azure Compute Disks.

https://learn.microsoft.com/en-us/rest/api/compute/disks/list?tabs=HTTP#disk

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|extended_location|`json`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|managed_by|`utf8`|
|managed_by_extended|`list<item: utf8, nullable>`|
|name|`utf8`|
|type|`utf8`|