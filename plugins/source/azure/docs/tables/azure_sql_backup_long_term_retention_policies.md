# Table: azure_sql_backup_long_term_retention_policies



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
|sql_database_id|String|
|weekly_retention|String|
|monthly_retention|String|
|yearly_retention|String|
|week_of_year|Int|
|id (PK)|String|
|name|String|
|type|String|