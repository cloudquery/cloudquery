# Table: azure_sql_encryption_protectors


The primary key for this table is **id**.

## Relations
This table depends on [`azure_sql_servers`](azure_sql_servers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|sql_server_id|String|
|kind|String|
|location|String|
|subregion|String|
|server_key_name|String|
|server_key_type|String|
|uri|String|
|thumbprint|String|
|id (PK)|String|
|name|String|
|type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|