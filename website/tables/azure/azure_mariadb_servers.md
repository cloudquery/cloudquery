# Table: azure_mariadb_servers

This table shows data for Azure MariaDB Servers.

https://learn.microsoft.com/en-us/rest/api/mariadb/servers/list?tabs=HTTP#server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_mariadb_servers:
  - [azure_mariadb_server_configurations](azure_mariadb_server_configurations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Geo-redundant backup should be enabled for Azure Database for MariaDB

```sql
SELECT
  'Geo-redundant backup should be enabled for Azure Database for MariaDB'
    AS title,
  subscription_id,
  id,
  CASE
  WHEN properties->'storageProfile'->>'geoRedundantBackup'
  IS DISTINCT FROM 'Enabled'
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_mariadb_servers;
```


