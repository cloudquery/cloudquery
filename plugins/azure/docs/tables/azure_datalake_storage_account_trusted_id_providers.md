
# Table: azure_datalake_storage_account_trusted_id_providers
TrustedIDProvider data Lake Store trusted identity provider information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|storage_account_cq_id|uuid|Unique CloudQuery ID of azure_datalake_storage_accounts table (FK)|
|id_provider|text|The URL of this trusted identity provider|
|id|text|The resource identifier|
|name|text|The resource name|
|type|text|The resource type|
