# Table: azure_compute_disks

This table shows data for Azure Compute Disks.

https://learn.microsoft.com/en-us/rest/api/compute/disks/list?tabs=HTTP#disk

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
|extended_location|JSON|
|properties|JSON|
|sku|JSON|
|tags|JSON|
|zones|StringArray|
|id (PK)|String|
|managed_by|String|
|managed_by_extended|StringArray|
|name|String|
|type|String|