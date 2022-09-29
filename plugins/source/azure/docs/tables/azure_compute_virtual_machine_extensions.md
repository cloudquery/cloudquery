# Table: azure_compute_virtual_machine_extensions


The primary key for this table is **id**.

## Relations
This table depends on [`azure_compute_virtual_machines`](azure_compute_virtual_machines.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|compute_virtual_machine_id|String|
|force_update_tag|String|
|publisher|String|
|type_handler_version|String|
|auto_upgrade_minor_version|Bool|
|enable_automatic_upgrade|Bool|
|provisioning_state|String|
|instance_view|JSON|
|id (PK)|String|
|name|String|
|location|String|
|tags|JSON|
|type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|