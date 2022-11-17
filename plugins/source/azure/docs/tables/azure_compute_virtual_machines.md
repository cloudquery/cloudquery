# Table: azure_compute_virtual_machines

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4#VirtualMachine

The primary key for this table is **id**.

## Relations

The following tables depend on azure_compute_virtual_machines:
  - [azure_compute_virtual_machine_scale_sets](azure_compute_virtual_machine_scale_sets.md)

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
|application_profile|JSON|
|availability_set|JSON|
|billing_profile|JSON|
|capacity_reservation|JSON|
|diagnostics_profile|JSON|
|eviction_policy|String|
|extensions_time_budget|String|
|hardware_profile|JSON|
|host|JSON|
|host_group|JSON|
|license_type|String|
|network_profile|JSON|
|os_profile|JSON|
|platform_fault_domain|Int|
|priority|String|
|proximity_placement_group|JSON|
|scheduled_events_profile|JSON|
|security_profile|JSON|
|storage_profile|JSON|
|user_data|String|
|virtual_machine_scale_set|JSON|
|instance_view|JSON|
|provisioning_state|String|
|time_created|Timestamp|
|vm_id|String|
|tags|JSON|
|zones|StringArray|
|id (PK)|String|
|name|String|
|resources|JSON|
|type|String|