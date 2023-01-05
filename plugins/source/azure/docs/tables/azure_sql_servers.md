# Table: azure_sql_servers

The primary key for this table is **id**.

## Relations

The following tables depend on azure_sql_servers:
  - [azure_sql_server_vulnerability_assessments](azure_sql_server_vulnerability_assessments.md)
  - [azure_sql_server_admins](azure_sql_server_admins.md)
  - [azure_sql_encryption_protectors](azure_sql_encryption_protectors.md)
  - [azure_sql_databases](azure_sql_databases.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|identity|JSON|
|properties|JSON|
|tags|JSON|
|id (PK)|String|
|kind|String|
|name|String|
|type|String|