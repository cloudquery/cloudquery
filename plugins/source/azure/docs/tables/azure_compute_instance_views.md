# Table: azure_compute_instance_views

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute#VirtualMachineInstanceView

The primary key for this table is **_cq_id**.

## Relations
This table depends on [azure_compute_virtual_machines](azure_compute_virtual_machines.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|compute_virtual_machine_id|String|
|platform_update_domain|Int|
|platform_fault_domain|Int|
|computer_name|String|
|os_name|String|
|os_version|String|
|hyper_v_generation|String|
|rdp_thumb_print|String|
|vm_agent|JSON|
|maintenance_redeploy_status|JSON|
|disks|JSON|
|extensions|JSON|
|vm_health|JSON|
|boot_diagnostics|JSON|
|assigned_host|String|
|statuses|JSON|
|patch_status|JSON|