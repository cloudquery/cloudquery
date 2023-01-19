# Table: azure_compute_virtual_machine_scale_sets

https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-scale-sets/list?tabs=HTTP#virtualmachinescaleset

The primary key for this table is **id**.

## Relations

The following tables depend on azure_compute_virtual_machine_scale_sets:
  - [azure_compute_virtual_machine_scale_set_vms](azure_compute_virtual_machine_scale_set_vms.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|location|String|
|extended_location|JSON|
|identity|JSON|
|plan|JSON|
|properties|JSON|
|sku|JSON|
|tags|JSON|
|zones|StringArray|
|name|String|
|type|String|