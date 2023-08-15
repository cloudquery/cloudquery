# Table: azure_sql_server_advanced_threat_protection_settings

This table shows data for Azure SQL Server Advanced Threat Protection Settings.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/server-advanced-threat-protection-settings/list-by-server?tabs=HTTP#advancedthreatprotectionstate

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
|system_data|`json`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that Advanced Threat Protection (ATP) on a SQL server is set to "Enabled" (Automated)

```sql
SELECT
  'Ensure that Advanced Threat Protection (ATP) on a SQL server is set to "Enabled" (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN atp.properties->>'state' IS DISTINCT FROM 'Enabled' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  JOIN azure_sql_server_advanced_threat_protection_settings AS atp ON
      s._cq_id = atp._cq_parent_id;
```


