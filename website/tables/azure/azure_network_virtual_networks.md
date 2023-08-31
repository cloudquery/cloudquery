# Table: azure_network_virtual_networks

This table shows data for Azure Network Virtual Networks.

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/virtual-networks/list-all?tabs=HTTP#virtualnetwork

The primary key for this table is **id**.

## Relations

The following tables depend on azure_network_virtual_networks:
  - [azure_network_virtual_network_subnets](azure_network_virtual_network_subnets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|extended_location|`json`|
|id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|etag|`utf8`|
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


