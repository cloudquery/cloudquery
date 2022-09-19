
# Table: azure_keyvault_managed_hsm
Managed HSM resource information with extended details.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|tenant_id|uuid|The Azure Active Directory tenant ID that should be used for authenticating requests to the managed HSM pool.|
|initial_admin_object_ids|text[]|Array of initial administrators object ids for this managed hsm pool.|
|hsm_uri|text|The URI of the managed hsm pool for performing operations on keys.|
|enable_soft_delete|boolean|Property to specify whether the 'soft delete' functionality is enabled for this managed HSM pool|
|soft_delete_retention_in_days|integer|Soft delete data retention days.|
|enable_purge_protection|boolean|Property specifying whether protection against purge is enabled for this managed HSM pool.|
|create_mode|text|The create mode to indicate whether the resource is being created or is being recovered from a deleted resource.|
|status_message|text|Resource Status Message.|
|provisioning_state|text|Provisioning state|
|id|text|The Azure Resource Manager resource ID for the managed HSM Pool.|
|name|text|The name of the managed HSM Pool.|
|type|text|The resource type of the managed HSM Pool.|
|location|text|The supported Azure location where the managed HSM Pool should be created.|
|sku_family|text|SKU Family of the managed HSM Pool.|
|sku_name|text|SKU of the managed HSM Pool.|
|tags|jsonb|Resource tags|
