# Table: azure_compute_virtual_machine_extensions

This table shows data for Azure Compute Virtual Machine Extensions.

https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-extensions/list?tabs=HTTP#virtualmachineextension

The primary key for this table is **id**.

## Relations

This table depends on [azure_compute_virtual_machines](azure_compute_virtual_machines).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|