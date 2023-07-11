# Table: azure_mysql_servers

This table shows data for Azure MySQL Servers.

https://learn.microsoft.com/en-us/rest/api/mysql/singleserver/servers(2017-12-01)/list?tabs=HTTP#server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_mysql_servers:
  - [azure_mysql_server_configurations](azure_mysql_server_configurations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|identity|`json`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Geo-redundant backup should be enabled for Azure Database for MySQL

```sql
SELECT
  'Geo-redundant backup should be enabled for Azure Database for MySQL'
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
  azure_mysql_servers;
```

### Ensure "Enforce SSL connection" is set to "ENABLED" for MySQL Database Server (Automated)

```sql
SELECT
  'Ensure "Enforce SSL connection" is set to "ENABLED" for MySQL Database Server (Automated)'
    AS title,
  subscription_id,
  id AS server_id,
  CASE
  WHEN properties->>'sslEnforcement' IS DISTINCT FROM 'Enabled' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_mysql_servers;
```


