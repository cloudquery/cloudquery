
# Table: azure_keyvault_vault_access_policies
AccessPolicyEntry an identity that have access to the key vault All identities in the array must use the same tenant ID as the key vault's tenant ID
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vault_cq_id|uuid|Unique CloudQuery ID of azure_keyvault_vaults table (FK)|
|tenant_id|uuid|The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault|
|object_id|text|The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault The object ID must be unique for the list of access policies|
|application_id|uuid|Application ID of the client making request on behalf of a principal|
|permissions_keys|text[]|Permissions to keys|
|permissions_secrets|text[]|Permissions to secrets|
|permissions_certificates|text[]|Permissions to certificates|
|permissions_storage|text[]|Permissions to storage accounts|
