# Table: azure_sql_server_administrators

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql#ServerAzureADAdministrator

The primary key for this table is **id**.

## Relations
This table depends on [azure_sql_servers](azure_sql_servers.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|administrator_type|String|
|login|String|
|sid|String|
|tenant_id|String|
|azure_ad_only_authentication|Bool|
|id (PK)|String|
|name|String|
|type|String|
|server_id|String|