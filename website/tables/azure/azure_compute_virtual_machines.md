# Table: azure_compute_virtual_machines

This table shows data for Azure Compute Virtual Machines.

https://learn.microsoft.com/en-us/rest/api/compute/virtual-machines/list?tabs=HTTP#virtualmachine

The primary key for this table is **id**.

## Relations

The following tables depend on azure_compute_virtual_machines:
  - [azure_compute_virtual_machine_extensions](azure_compute_virtual_machine_extensions)
  - [azure_compute_virtual_machine_patch_assessments](azure_compute_virtual_machine_patch_assessments)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|subscription_id|utf8|
|instance_view|json|
|location|utf8|
|extended_location|json|
|identity|json|
|plan|json|
|properties|json|
|tags|json|
|zones|list<item: utf8, nullable>|
|id (PK)|utf8|
|name|utf8|
|resources|json|
|type|utf8|