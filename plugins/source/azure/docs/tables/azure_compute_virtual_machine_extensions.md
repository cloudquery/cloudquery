# Table: azure_compute_virtual_machine_extensions

The primary key for this table is **id**.

## Relations

This table depends on [azure_compute_virtual_machines](azure_compute_virtual_machines.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|location|String|
|properties|JSON|
|tags|JSON|
|name|String|
|type|String|