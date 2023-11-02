# Table: azure_compute_snapshots

This table shows data for Azure Compute Snapshots.

https://learn.microsoft.com/en-us/rest/api/compute/snapshots/list?tabs=HTTP#snapshot

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|extended_location|`json`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|managed_by|`utf8`|
|name|`utf8`|
|type|`utf8`|