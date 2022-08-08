
# Table: azure_compute_virtual_machine_resources
VirtualMachineExtension describes a Virtual Machine Extension.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_cq_id|uuid|Unique CloudQuery ID of azure_compute_virtual_machines table (FK)|
|virtual_machine_id|text|ID of azure_compute_virtual_machines table (FK)|
|force_update_tag|text|How the extension handler should be forced to update even if the extension configuration has not changed.|
|publisher|text|The name of the extension handler publisher.|
|type_handler_version|text|Specifies the version of the script handler.|
|auto_upgrade_minor_version|boolean|Indicates whether the extension should use a newer minor version if one is available at deployment time|
|enable_automatic_upgrade|boolean|Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.|
|settings|jsonb|Json formatted public settings for the extension.|
|protected_settings|jsonb|The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.|
|extension_type|text|Type of the extension|
|provisioning_state|text|The provisioning state, which only appears in the response.|
|instance_view|jsonb|The virtual machine extension instance view.|
|id|text|Resource Id|
|name|text|Resource name|
|type|text|Resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
