
# Table: azure_compute_virtual_machine_scale_sets
VirtualMachineScaleSet describes a Virtual Machine Scale Set
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|sku_name|text|The sku name|
|sku_tier|text|Specifies the tier of virtual machines in a scale set<br /><br /> Possible Values:<br /><br /> **Standard**<br /><br /> **Basic**|
|sku_capacity|bigint|Specifies the number of virtual machines in the scale set|
|plan_name|text|The plan ID|
|plan_publisher|text|The publisher ID|
|plan_product|text|Specifies the product of the image from the marketplace|
|plan_promotion_code|text|The promotion code|
|upgrade_policy|jsonb|The upgrade policy|
|automatic_repairs_policy_enabled|boolean|Specifies whether automatic repairs should be enabled on the virtual machine scale set|
|automatic_repairs_policy_grace_period|text|The amount of time for which automatic repairs are suspended due to a state change on VM|
|os_profile_computer_name_prefix|text|Specifies the computer name prefix for all of the virtual machines in the scale set|
|os_profile_admin_username|text|Specifies the name of the administrator account|
|os_profile_admin_password|text|Specifies the password of the administrator account|
|os_profile_custom_data|text|Specifies a base-64 encoded string of custom data|
|os_profile_windows_configuration|jsonb|Specifies Windows operating system settings on the virtual machine|
|os_profile_linux_configuration|jsonb|Specifies the Linux operating system settings on the virtual machine|
|storage_profile|jsonb|Specifies the storage settings for the virtual machine disks|
|network_profile|jsonb|Specifies properties of the network interfaces of the virtual machines in the scale set|
|security_profile|jsonb|Specifies the Security related profile settings for the virtual machines in the scale set|
|diagnostics_profile|jsonb|Specifies the boot diagnostic settings state|
|extension_profile_extensions_time_budget|text|Specifies the time alloted for all extensions to start|
|license_type|text|Specifies that the image or disk that is being used was licensed on-premises|
|priority|text|Specifies the priority for the virtual machines in the scale set|
|eviction_policy|text|Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set|
|billing_profile_max_price|float|Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS|
|scheduled_events_profile|jsonb|Specifies Scheduled Event related configurations|
|user_data|text|UserData for the virtual machines in the scale set, which must be base-64 encoded|
|provisioning_state|text|The provisioning state, which only appears in the response|
|overprovision|boolean|Specifies whether the Virtual Machine Scale Set should be overprovisioned|
|do_not_run_extensions_on_overprovisioned_vms|boolean|When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept|
|unique_id|text|Specifies the ID which uniquely identifies a Virtual Machine Scale Set|
|single_placement_group|boolean|When true this limits the scale set to a single placement group, of max size 100 virtual machines|
|zone_balance|boolean|Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage|
|platform_fault_domain_count|integer|Fault Domain count for each placement group|
|proximity_placement_group_id|text|Proximity placement group resource id|
|host_group_id|text|Host group resource id|
|additional_capabilities_ultra_ssd_enabled|boolean|The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS storage account type on the VM or VMSS|
|scale_in_policy_rules|text[]|The rules to be followed when scaling-in a virtual machine scale set|
|orchestration_mode|text|Specifies the orchestration mode for the virtual machine scale set|
|identity_principal_id|text|The principal id of virtual machine scale set identity|
|identity_tenant_id|text|The tenant id associated with the virtual machine scale set|
|identity_type|text|The type of identity used for the virtual machine scale set|
|identity_user_assigned_identities|jsonb|The list of user identities associated with the virtual machine scale set|
|zones|text[]|The virtual machine scale set zones|
|extended_location_name|text|The name of the extended location|
|extended_location_type|text|The type of the extended location|
|id|text|Resource Id|
|name|text|Resource name|
|type|text|Resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
