# Table: azure_container_replications



The primary key for this table is **id**.

## Relations
This table depends on [azure_container_registries](azure_container_registries.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|container_registry_id|String|
|provisioning_state|String|
|status|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|