
# Table: azure_compute_disk_encryption_settings
Azure compute disk encryption setting
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|disk_id|uuid|Unique ID of azure_compute_disks table (FK)|
|disk_encryption_key_source_vault_id|text|Resource Id|
|disk_encryption_key_secret_url|text|Url pointing to a key or secret in KeyVault|
|key_encryption_key_source_vault_id|text|Resource Id|
|key_encryption_key_key_url|text|Url pointing to a key or secret in KeyVault|
