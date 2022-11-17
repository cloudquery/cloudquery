# Table: azure_compute_virtual_machine_scale_sets

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4#VirtualMachineScaleSet

The primary key for this table is **id**.

## Relations
This table depends on [azure_compute_virtual_machines](azure_compute_virtual_machines.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|extended_location|JSON|
|identity|JSON|
|plan|JSON|
|additional_capabilities|JSON|
|automatic_repairs_policy|JSON|
|do_not_run_extensions_on_overprovisioned_vms|Bool|
|host_group|JSON|
|orchestration_mode|String|
|overprovision|Bool|
|platform_fault_domain_count|Int|
|priority_mix_policy|JSON|
|proximity_placement_group|JSON|
|scale_in_policy|JSON|
|single_placement_group|Bool|
|spot_restore_policy|JSON|
|upgrade_policy|JSON|
|virtual_machine_profile|JSON|
|zone_balance|Bool|
|provisioning_state|String|
|time_created|Timestamp|
|unique_id|String|
|sku|JSON|
|tags|JSON|
|zones|StringArray|
|id (PK)|String|
|name|String|
|type|String|
|virtual_machine_id|String|