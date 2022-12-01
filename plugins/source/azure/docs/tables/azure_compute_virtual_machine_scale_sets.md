# Table: azure_compute_virtual_machine_scale_sets

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute#VirtualMachineScaleSet

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sku|JSON|
|plan|JSON|
|upgrade_policy|JSON|
|automatic_repairs_policy|JSON|
|virtual_machine_profile|JSON|
|provisioning_state|String|
|overprovision|Bool|
|unique_id|String|
|single_placement_group|Bool|
|zone_balance|Bool|
|platform_fault_domain_count|Int|
|proximity_placement_group|JSON|
|host_group|JSON|
|additional_capabilities|JSON|
|scale_in_policy|JSON|
|orchestration_mode|String|
|identity|JSON|
|zones|StringArray|
|extended_location|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|
|do_not_run_extensions_on_overprovisioned_vms|Bool|