# Table: azure_compute_virtual_machines

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute#VirtualMachine

The primary key for this table is **id**.

## Relations

The following tables depend on azure_compute_virtual_machines:
  - [azure_compute_instance_views](azure_compute_instance_views.md)
  - [azure_compute_virtual_machine_extensions](azure_compute_virtual_machine_extensions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|plan|JSON|
|hardware_profile|JSON|
|storage_profile|JSON|
|additional_capabilities|JSON|
|os_profile|JSON|
|network_profile|JSON|
|security_profile|JSON|
|diagnostics_profile|JSON|
|availability_set|JSON|
|virtual_machine_scale_set|JSON|
|proximity_placement_group|JSON|
|priority|String|
|eviction_policy|String|
|billing_profile|JSON|
|host|JSON|
|host_group|JSON|
|provisioning_state|String|
|instance_view|JSON|
|license_type|String|
|vm_id|String|
|extensions_time_budget|String|
|platform_fault_domain|Int|
|scheduled_events_profile|JSON|
|user_data|String|
|resources|JSON|
|identity|JSON|
|zones|StringArray|
|extended_location|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|