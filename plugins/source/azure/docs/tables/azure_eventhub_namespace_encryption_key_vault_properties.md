
# Table: azure_eventhub_namespace_encryption_key_vault_properties
KeyVaultProperties properties to configure keyVault Properties
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|namespace_cq_id|uuid|Unique CloudQuery ID of azure_eventhub_namespaces table (FK)|
|key_name|text|Name of the Key from KeyVault|
|key_vault_uri|text|Uri of KeyVault|
|key_version|text|Key Version|
