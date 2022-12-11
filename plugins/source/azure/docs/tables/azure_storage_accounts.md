# Table: azure_storage_accounts



The primary key for this table is **id**.

## Relations

The following tables depend on azure_storage_accounts:
  - [azure_storage_tables](azure_storage_tables.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|location|String|
|extended_location|JSON|
|identity|JSON|
|properties|JSON|
|tags|JSON|
|id (PK)|String|
|kind|String|
|name|String|
|sku|JSON|
|type|String|