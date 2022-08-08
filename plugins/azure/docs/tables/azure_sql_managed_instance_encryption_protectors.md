
# Table: azure_sql_managed_instance_encryption_protectors
ManagedInstanceEncryptionProtector the managed instance encryption protector
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|managed_instance_cq_id|uuid|Unique CloudQuery ID of azure_sql_managed_instances table (FK)|
|kind|text|Kind of encryption protector|
|server_key_name|text|The name of the managed instance key|
|server_key_type|text|The encryption protector type like 'ServiceManaged', 'AzureKeyVault'|
|uri|text|The URI of the server key|
|thumbprint|text|Thumbprint of the server key|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
