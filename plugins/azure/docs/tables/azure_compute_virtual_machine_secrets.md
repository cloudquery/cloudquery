
# Table: azure_compute_virtual_machine_secrets
VaultSecretGroup describes a set of certificates which are all in the same Key Vault
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_cq_id|uuid|Unique CloudQuery ID of azure_compute_virtual_machines table (FK)|
|virtual_machine_id|text|ID of azure_compute_virtual_machines table (FK)|
|source_vault_id|text|Source vault Id|
