# Table: azure_sql_server_virtual_network_rules

This table shows data for Azure SQL Server Virtual Network Rules.

https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/virtual-network-rules/list-by-server?tabs=HTTP#virtualnetworkrule

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

### SQL Server should use a virtual network service endpoint

```sql
WITH
  subs
    AS (
      SELECT
        subscription_id,
        jsonb_array_elements(properties->'subnets') AS subnet,
        properties->>'provisioningState' AS provisioning_state
      FROM
        azure_network_virtual_networks
    ),
  secured_servers
    AS (
      SELECT
        s._cq_id
      FROM
        azure_sql_servers AS s
        LEFT JOIN azure_sql_server_virtual_network_rules AS r ON
            s._cq_id = r._cq_parent_id
        LEFT JOIN subs ON
            r.properties->>'virtualNetworkSubnetId' = subs.subnet->>'id'
      WHERE
        (r.properties->'virtualNetworkSubnetId') IS NOT NULL
        AND subs.provisioning_state = 'Succeeded'
    )
SELECT
  'SQL Server should use a virtual network service endpoint' AS title,
  subscription_id,
  id,
  CASE
  WHEN ss._cq_id IS NULL THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s LEFT JOIN secured_servers AS ss ON s._cq_id = ss._cq_id;
```


