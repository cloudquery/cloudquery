# Table: azure_sql_transparent_data_encryptions

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql#LogicalDatabaseTransparentDataEncryption

The primary key for this table is **id**.

## Relations
This table depends on [azure_sql_databases](azure_sql_databases.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|state|String|
|id (PK)|String|
|name|String|
|type|String|
|database_id|String|