# Table: azure_compute_virtual_machine_scale_set_network_interfaces

This table shows data for Azure Compute Virtual Machine Scale Set Network Interfaces.

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/network-interface-in-vm-ss/list-virtual-machine-scale-set-network-interfaces?tabs=HTTP#networkinterface

The primary key for this table is **id**.

## Relations

This table depends on [azure_compute_virtual_machine_scale_sets](azure_compute_virtual_machine_scale_sets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|extended_location|JSON|
|id (PK)|String|
|location|String|
|properties|JSON|
|tags|JSON|
|etag|String|
|name|String|
|type|String|