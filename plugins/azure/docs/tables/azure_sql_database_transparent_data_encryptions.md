
# Table: azure_sql_database_transparent_data_encryptions
Azure sql database encryption
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_id|uuid|Unique ID of azure_sql_databases table (FK)|
|location|text|Resource location|
|status|text|The status of the database transparent data encryption Possible values include: 'TransparentDataEncryptionStatusEnabled', 'TransparentDataEncryptionStatusDisabled'|
|resource_id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
