
# Table: azure_compute_virtual_machines
VirtualMachine describes a Virtual Machine.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|plan_name|text|The plan ID.|
|plan_publisher|text|The publisher ID.|
|plan_product|text|Specifies the product of the image from the marketplace|
|plan_promotion_code|text|The promotion code.|
|hardware_profile_vm_size|text|Specifies the size of the virtual machine|
|storage_profile|jsonb|Specifies the storage settings for the virtual machine disks.|
|additional_capabilities_ultra_ssd_enabled|boolean|The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS storage account type on the VM or VMSS|
|computer_name|text|Specifies the host OS name of the virtual machine|
|admin_username|text|Specifies the name of the administrator account|
|admin_password|text|Specifies the password of the administrator account|
|custom_data|text|Specifies a base-64 encoded string of custom data|
|windows_configuration_provision_vm_agent|boolean|Indicates whether virtual machine agent should be provisioned on the virtual machine|
|windows_configuration_enable_automatic_updates|boolean|Indicates whether Automatic Updates is enabled for the Windows virtual machine|
|windows_configuration_time_zone|text|Specifies the time zone of the virtual machine|
|windows_configuration_additional_unattend_content|jsonb|Specifies additional base-64 encoded XML formatted information that can be included in the Unattend.xml file, which is used by Windows Setup.|
|windows_configuration_patch_settings_patch_mode|text|You  control the application of patches to a virtual machine|
|windows_configuration_patch_settings_enable_hotpatching|boolean|Enables customers to patch their Azure VMs without requiring a reboot|
|windows_configuration_patch_settings_assessment_mode|text|The platform will trigger periodic patch assessments|
|linux_configuration_disable_password_authentication|boolean|Specifies whether password authentication should be disabled.|
|linux_configuration_ssh_public_keys|jsonb|The list of SSH public keys used to authenticate with linux based VMs.|
|linux_configuration_provision_vm_agent|boolean|Indicates whether virtual machine agent should be provisioned on the virtual machine|
|linux_configuration_patch_settings_patch_mode|text|The virtual machine's default patching configuration is used|
|linux_configuration_patch_settings_assessment_mode|text|You control the timing of patch assessments on a virtual machine|
|allow_extension_operations|boolean|Specifies whether extension operations should be allowed on the virtual machine|
|require_guest_provision_signal|boolean|Specifies whether the guest provision signal is required to infer provision success of the virtual machine|
|network_profile_network_interfaces|jsonb|Specifies the list of resource Ids for the network interfaces associated with the virtual machine.|
|network_profile_network_api_version|text|specifies the Microsoft.Network API version used when creating networking resources in the Network Interface Configurations|
|network_profile_network_interface_configurations|jsonb|Specifies the networking configurations that will be used to create the virtual machine networking resources.|
|security_profile_uefi_settings_secure_boot_enabled|boolean|Specifies whether secure boot should be enabled on the virtual machine|
|security_profile_uefi_settings_v_tpm_enabled|boolean|Specifies whether vTPM should be enabled on the virtual machine|
|security_profile_encryption_at_host|boolean|This property can be used by user in the request to enable or disable the Host Encryption for the virtual machine or virtual machine scale set|
|security_profile_security_type|text|Specifies the SecurityType of the virtual machine|
|diagnostics_profile_boot_diagnostics_enabled|boolean|Whether boot diagnostics should be enabled on the Virtual Machine.|
|diagnostics_profile_boot_diagnostics_storage_uri|text|Uri of the storage account to use for placing the console output and screenshot|
|availability_set_id|text|Resource Id|
|virtual_machine_scale_set_id|text|Resource Id|
|proximity_placement_group_id|text|Resource Id|
|priority|text|Specifies the priority for the virtual machine|
|eviction_policy|text|Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set|
|billing_profile_max_price|float|Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS|
|host_id|text|Resource Id|
|host_group_id|text|Resource Id|
|provisioning_state|text|The provisioning state, which only appears in the response.|
|instance_view|jsonb|The virtual machine instance view.|
|license_type|text|Specifies that the image or disk that is being used was licensed on-premises|
|vm_id|text|Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.|
|extensions_time_budget|text|Specifies the time alloted for all extensions to start|
|platform_fault_domain|integer|Specifies the scale set logical fault domain into which the Virtual Machine will be created.|
|scheduled_events_profile|jsonb|Specifies Scheduled Event related configurations.|
|user_data|text|UserData for the VM, which must be base-64 encoded|
|identity_principal_id|text|The principal id of virtual machine identity|
|identity_tenant_id|text|The tenant id associated with the virtual machine|
|identity_type|text|The type of identity used for the virtual machine|
|identity_user_assigned_identities|jsonb|The list of user identities associated with the Virtual Machine|
|zones|text[]|The virtual machine zones.|
|extended_location_name|text|The name of the extended location.|
|extended_location_type|text|The type of the extended location|
|id|text|Resource Id|
|name|text|Resource name|
|type|text|Resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
