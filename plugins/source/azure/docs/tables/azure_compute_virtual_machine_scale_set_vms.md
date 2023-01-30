# Table: azure_compute_virtual_machine_scale_set_vms

https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-scale-set-vms/list?tabs=HTTP#virtualmachinescalesetvm

The primary key for this table is **id**.

## Relations

This table depends on [azure_compute_virtual_machine_scale_sets](azure_compute_virtual_machine_scale_sets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|identity|JSON|
|plan|JSON|
|properties|JSON|
|tags|JSON|
|id (PK)|String|
|instance_id|String|
|name|String|
|resources|JSON|
|sku|JSON|
|type|String|
|zones|StringArray|