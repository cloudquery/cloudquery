# Table: azure_compute_virtual_machines

https://learn.microsoft.com/en-us/rest/api/compute/virtual-machines/list?tabs=HTTP#virtualmachine

The primary key for this table is **id**.

## Relations

The following tables depend on azure_compute_virtual_machines:
  - [azure_compute_virtual_machine_extensions](azure_compute_virtual_machine_extensions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|instance_view|JSON|
|id (PK)|String|
|location|String|
|extended_location|JSON|
|identity|JSON|
|plan|JSON|
|properties|JSON|
|tags|JSON|
|zones|StringArray|
|name|String|
|resources|JSON|
|type|String|