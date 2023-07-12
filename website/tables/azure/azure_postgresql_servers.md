# Table: azure_postgresql_servers

This table shows data for Azure PostgreSQL Servers.

https://learn.microsoft.com/en-us/rest/api/postgresql/singleserver/servers/list?tabs=HTTP#server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_postgresql_servers:
  - [azure_postgresql_databases](azure_postgresql_databases)
  - [azure_postgresql_server_configurations](azure_postgresql_server_configurations)
  - [azure_postgresql_server_firewall_rules](azure_postgresql_server_firewall_rules)

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

### Ensure server parameter "connection_throttling" is set to "ON" for PostgreSQL Database Server (Automated)

```sql
WITH
  value_check
    AS (
      SELECT
        aps._cq_id, apsc.properties->>'value' AS value
      FROM
        azure_postgresql_servers AS aps
        LEFT JOIN azure_postgresql_server_configurations AS apsc ON
            aps._cq_id = apsc._cq_parent_id
      WHERE
        apsc.name = 'connection_throttling'
    )
SELECT
  'Ensure server parameter "connection_throttling" is set to "ON" for PostgreSQL Database Server (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN v.value IS NULL OR v.value != 'on' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_postgresql_servers AS s
  LEFT JOIN value_check AS v ON s._cq_id = v._cq_id;
```

### Ensure server parameter "log_checkpoints" is set to "ON" for PostgreSQL Database Server (Automated)

```sql
WITH
  value_check
    AS (
      SELECT
        aps._cq_id, apsc.properties->>'value' AS value
      FROM
        azure_postgresql_servers AS aps
        LEFT JOIN azure_postgresql_server_configurations AS apsc ON
            aps._cq_id = apsc._cq_parent_id
      WHERE
        apsc.name = 'log_checkpoints'
    )
SELECT
  'Ensure server parameter "log_checkpoints" is set to "ON" for PostgreSQL Database Server (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN v.value IS NULL OR v.value != 'on' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_postgresql_servers AS s
  LEFT JOIN value_check AS v ON s._cq_id = v._cq_id;
```

### Ensure server parameter "log_connections" is set to "ON" for PostgreSQL Database Server (Automated)

```sql
WITH
  value_check
    AS (
      SELECT
        aps._cq_id, apsc.properties->>'value' AS value
      FROM
        azure_postgresql_servers AS aps
        LEFT JOIN azure_postgresql_server_configurations AS apsc ON
            aps._cq_id = apsc._cq_parent_id
      WHERE
        apsc.name = 'log_connections'
    )
SELECT
  'Ensure server parameter "log_connections" is set to "ON" for PostgreSQL Database Server (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN v.value IS NULL OR v.value != 'on' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_postgresql_servers AS s
  LEFT JOIN value_check AS v ON s._cq_id = v._cq_id;
```

### Ensure server parameter "log_disconnections" is set to "ON" for PostgreSQL Database Server (Automated)

```sql
WITH
  value_check
    AS (
      SELECT
        aps._cq_id, apsc.properties->>'value' AS value
      FROM
        azure_postgresql_servers AS aps
        LEFT JOIN azure_postgresql_server_configurations AS apsc ON
            aps._cq_id = apsc._cq_parent_id
      WHERE
        apsc.name = 'log_disconnections'
    )
SELECT
  'Ensure server parameter "log_disconnections" is set to "ON" for PostgreSQL Database Server (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN v.value IS NULL OR v.value != 'on' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_postgresql_servers AS s
  LEFT JOIN value_check AS v ON s._cq_id = v._cq_id;
```

### Ensure server parameter "log_retention_days" is greater than 3 days for PostgreSQL Database Server (Automated)

```sql
WITH
  value_check
    AS (
      SELECT
        aps._cq_id, apsc.properties->>'value' AS value
      FROM
        azure_postgresql_servers AS aps
        LEFT JOIN azure_postgresql_server_configurations AS apsc ON
            aps._cq_id = apsc._cq_parent_id
      WHERE
        apsc.name = 'log_retention_days'
    )
SELECT
  'Ensure server parameter "log_retention_days" is greater than 3 days for PostgreSQL Database Server (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN v.value IS NULL OR v.value::INT8 < 3 THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_postgresql_servers AS s
  LEFT JOIN value_check AS v ON s._cq_id = v._cq_id;
```

### Geo-redundant backup should be enabled for Azure Database for PostgreSQL

```sql
SELECT
  'Geo-redundant backup should be enabled for Azure Database for PostgreSQL'
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
  azure_postgresql_servers;
```

### Ensure "Enforce SSL connection" is set to "ENABLED" for PostgreSQL Database Server (Automated)

```sql
SELECT
  'Ensure "Enforce SSL connection" is set to "ENABLED" for PostgreSQL Database Server (Automated)'
    AS title,
  subscription_id,
  id AS server_id,
  CASE
  WHEN properties->>'sslEnforcement' IS DISTINCT FROM 'Enabled' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_postgresql_servers;
```


