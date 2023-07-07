# Table: azure_sql_server_database_long_term_retention_policies

This table shows data for Azure SQL Server Database Long Term Retention Policies.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/long-term-retention-policies/list-by-database?tabs=HTTP#longtermretentionpolicy

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
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Long-term geo-redundant backup should be enabled for Azure SQL Databases

```sql
SELECT
  'Long-term geo-redundant backup should be enabled for Azure SQL Databases'
    AS title,
  s.subscription_id,
  rp.id,
  CASE
  WHEN rp.id IS NULL
  OR (
      rp.properties->>'weeklyRetention' IS NOT DISTINCT FROM 'PT0S'
      AND rp.properties->>'monthlyRetention' IS NOT DISTINCT FROM 'PT0S'
      AND rp.properties->>'yearlyRetention' IS NOT DISTINCT FROM 'PT0S'
    )
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_databases AS asd ON s._cq_id = asd._cq_parent_id
  LEFT JOIN azure_sql_server_database_long_term_retention_policies AS rp ON
      asd._cq_id = rp._cq_parent_id;
```


