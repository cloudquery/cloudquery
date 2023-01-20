# Table: azure_sql_server_databases

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/databases/list-by-server?tabs=HTTP#database

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_servers](azure_sql_servers.md).

The following tables depend on azure_sql_server_databases:
  - [azure_sql_server_database_blob_auditing_policies](azure_sql_server_database_blob_auditing_policies.md)
  - [azure_sql_server_database_long_term_retention_policies](azure_sql_server_database_long_term_retention_policies.md)
  - [azure_sql_server_database_threat_protections](azure_sql_server_database_threat_protections.md)
  - [azure_sql_server_database_vulnerability_assessments](azure_sql_server_database_vulnerability_assessments.md)
  - [azure_sql_transparent_data_encryptions](azure_sql_transparent_data_encryptions.md)

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
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|kind|String|
|managed_by|String|
|name|String|
|type|String|