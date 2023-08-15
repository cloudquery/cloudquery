# Table: azure_sql_server_firewall_rules

This table shows data for Azure SQL Server Firewall Rules.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/firewall-rules/list-by-server?tabs=HTTP#firewallrule

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_servers](azure_sql_servers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|name|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure no SQL Databases allow ingress 0.0.0.0/0 (ANY IP)

```sql
SELECT
  'Ensure no SQL Databases allow ingress 0.0.0.0/0 (ANY IP)' AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN properties->>'startIpAddress' = '0.0.0.0'
  AND (
      properties->>'endIpAddress' = '0.0.0.0'
      OR properties->>'endIpAddress' = '255.255.255.255'
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_sql_server_firewall_rules;
```


