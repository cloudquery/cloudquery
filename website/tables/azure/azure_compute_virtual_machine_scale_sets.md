# Table: azure_compute_virtual_machine_scale_sets

This table shows data for Azure Compute Virtual Machine Scale Sets.

https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-scale-sets/list?tabs=HTTP#virtualmachinescaleset

The primary key for this table is **id**.

## Relations

The following tables depend on azure_compute_virtual_machine_scale_sets:
  - [azure_compute_virtual_machine_scale_set_network_interfaces](azure_compute_virtual_machine_scale_set_network_interfaces)
  - [azure_compute_virtual_machine_scale_set_vms](azure_compute_virtual_machine_scale_set_vms)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|extended_location|`json`|
|identity|`json`|
|plan|`json`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|