# Table: azure_compute_virtual_machine_extensions

https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-extensions/list?tabs=HTTP#virtualmachineextension

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
|subscription_id|String|
|location|String|
|properties|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|