# Table: azure_sql_transparent_data_encryptions

This table shows data for Azure SQL Transparent Data Encryptions.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/transparent-data-encryptions/list-by-database?tabs=HTTP#logicaldatabasetransparentdataencryption

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

### Ensure that "Data encryption" is set to "On" on a SQL Database (Automated)

```sql
SELECT
  'Ensure that "Data encryption" is set to "On" on a SQL Database (Automated)'
    AS title,
  s.subscription_id,
  asd.id AS database_id,
  CASE
  WHEN tde.properties->>'state' IS DISTINCT FROM 'Enabled' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_databases AS asd ON s._cq_id = asd._cq_parent_id
  LEFT JOIN azure_sql_transparent_data_encryptions AS tde ON
      asd._cq_id = tde._cq_parent_id
WHERE
  asd.name != 'master';
```


