# Table: azure_compute_virtual_machine_scale_set_vms

This table shows data for Azure Compute Virtual Machine Scale Set Vms.

https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-scale-set-vms/list?tabs=HTTP#virtualmachinescalesetvm

The primary key for this table is **id**.

## Relations

This table depends on [azure_compute_virtual_machine_scale_sets](azure_compute_virtual_machine_scale_sets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|identity|`json`|
|plan|`json`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|instance_id|`utf8`|
|name|`utf8`|
|resources|`json`|
|sku|`json`|
|type|`utf8`|
|zones|`list<item: utf8, nullable>`|