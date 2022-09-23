# Table: azure_sql_backup_long_term_retention_policies


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|sql_database_id|UUID|
|weekly_retention|String|
|monthly_retention|String|
|yearly_retention|String|
|week_of_year|Int|
|id (PK)|String|
|name|String|
|type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|