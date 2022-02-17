
# Table: azure_compute_virtual_machine_scale_set_extensions
VirtualMachineScaleSetExtension describes a Virtual Machine Scale Set Extension
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_scale_set_cq_id|uuid|Unique CloudQuery ID of azure_compute_virtual_machine_scale_sets table (FK)|
|type|text|The type of the resource|
|extension_type|text|The type of the extension|
|name|text|The name of the extension|
|force_update_tag|text|If a value is provided and is different from the previous value, the extension handler will be forced to update even if the extension configuration has not changed|
|publisher|text|The name of the extension handler publisher|
|type_handler_version|text|Specifies the version of the script handler|
|auto_upgrade_minor_version|boolean|Indicates whether the extension should use a newer minor version if one is available at deployment time|
|enable_automatic_upgrade|boolean|Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available|
|settings|jsonb|Json formatted public settings for the extension|
|protected_settings|jsonb|The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all|
|provisioning_state|text|The provisioning state, which only appears in the response|
|provision_after_extensions|text[]|Collection of extension names after which this extension needs to be provisioned|
|id|text|Resource Id|
