# Table: azure_sql_server_database_blob_auditing_policies

This table shows data for Azure SQL Server Database Blob Auditing Policies.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/database-blob-auditing-policies/list-by-database?tabs=HTTP#databaseblobauditingpolicy

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_server_databases](azure_sql_server_databases).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|kind|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that "Auditing" is set to "On" (Automated)

```sql
SELECT
  'Ensure that "Auditing" is set to "On" (Automated)' AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN assdbap.properties->>'state' != 'Enabled' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_databases AS assd ON s._cq_id = assd._cq_parent_id
  LEFT JOIN azure_sql_server_database_blob_auditing_policies AS assdbap ON
      assd._cq_id = assdbap._cq_parent_id;
```


