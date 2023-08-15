# Table: azure_sql_server_blob_auditing_policies

This table shows data for Azure SQL Server Blob Auditing Policies.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/server-blob-auditing-policies/list-by-server?tabs=HTTP#serverblobauditingpolicy

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_servers](azure_sql_servers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that "Auditing" Retention is "greater than 90 days" (Automated)

```sql
SELECT
  'Ensure that "Auditing" Retention is "greater than 90 days" (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN (assdbap.properties->'retentionDays')::INT8 < 90 THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_blob_auditing_policies AS assdbap ON
      s._cq_id = assdbap._cq_parent_id;
```

### Auditing on SQL server should be enabled

```sql
SELECT
  'Auditing on SQL server should be enabled' AS title,
  sub.id,
  sub.display_name AS subscription_name,
  CASE
  WHEN azure_sql_server_blob_auditing_policies._cq_parent_id
  = azure_sql_servers._cq_id
  AND sub.id = azure_sql_servers.subscription_id
  AND azure_sql_server_blob_auditing_policies.properties->>'state' = 'Disabled'
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_server_blob_auditing_policies,
  azure_sql_servers,
  azure_subscription_subscriptions AS sub;
```


