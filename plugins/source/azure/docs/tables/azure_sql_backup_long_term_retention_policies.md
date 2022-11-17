# Table: azure_sql_backup_long_term_retention_policies

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql#LongTermRetentionPolicy

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
|monthly_retention|String|
|week_of_year|Int|
|weekly_retention|String|
|yearly_retention|String|
|id (PK)|String|
|name|String|
|type|String|
|database_id|String|