# Table: azure_postgresql_server_firewall_rules

This table shows data for Azure PostgreSQL Server Firewall Rules.

https://learn.microsoft.com/en-us/rest/api/postgresql/singleserver/firewall-rules/list-by-server?tabs=HTTP#firewallrule

The primary key for this table is **id**.

## Relations

This table depends on [azure_postgresql_servers](azure_postgresql_servers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure "Allow access to Azure services" for PostgreSQL Database Server is disabled (Automated)

```sql
SELECT
  'Ensure "Allow access to Azure services" for PostgreSQL Database Server is disabled (Automated)'
    AS title,
  aps.subscription_id,
  aps.id AS server_id,
  CASE
  WHEN apsfr.name = 'AllowAllAzureIps'
  OR (
      apsfr.properties->>'startIPAddress' = '0.0.0.0'
      AND apsfr.properties->>'endIPAddress' = '0.0.0.0'
    )
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_postgresql_servers AS aps
  LEFT JOIN azure_postgresql_server_firewall_rules AS apsfr ON
      aps._cq_id = apsfr._cq_parent_id;
```


