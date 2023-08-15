# Table: azure_compute_virtual_machine_scale_set_network_interfaces

This table shows data for Azure Compute Virtual Machine Scale Set Network Interfaces.

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/network-interface-in-vm-ss/list-virtual-machine-scale-set-network-interfaces?tabs=HTTP#networkinterface

The primary key for this table is **id**.

## Relations

This table depends on [azure_compute_virtual_machine_scale_sets](azure_compute_virtual_machine_scale_sets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|extended_location|`json`|
|id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|