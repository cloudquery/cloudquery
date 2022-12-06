# Table: azure_sql_server_admins

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#ServerAzureADAdministrator

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
|sql_server_id|String|
|administrator_type|String|
|login|String|
|sid|UUID|
|tenant_id|UUID|
|azure_ad_only_authentication|Bool|
|id (PK)|String|
|name|String|
|type|String|