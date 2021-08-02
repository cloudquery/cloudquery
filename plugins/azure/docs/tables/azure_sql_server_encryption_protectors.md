
# Table: azure_sql_server_encryption_protectors
EncryptionProtector the server encryption protector
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_sql_servers table (FK)|
|kind|text|Kind of encryption protector This is metadata used for the Azure portal experience|
|location|text|Resource location|
|subregion|text|Subregion of the encryption protector|
|server_key_name|text|The name of the server key|
|server_key_type|text|The encryption protector type like 'ServiceManaged', 'AzureKeyVault' Possible values include: 'ServiceManaged', 'AzureKeyVault'|
|uri|text|The URI of the server key|
|thumbprint|text|Thumbprint of the server key|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
