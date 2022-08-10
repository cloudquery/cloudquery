
# Table: azure_compute_virtual_machine_scale_set_os_profile_secrets
VaultSecretGroup describes a set of certificates which are all in the same Key Vault
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_scale_set_cq_id|uuid|Unique CloudQuery ID of azure_compute_virtual_machine_scale_sets table (FK)|
|source_vault_id|text|Resource Id|
|vault_certificates|jsonb|The list of key vault references in SourceVault which contain certificates|
