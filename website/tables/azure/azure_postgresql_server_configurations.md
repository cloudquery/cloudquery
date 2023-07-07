# Table: azure_postgresql_server_configurations

This table shows data for Azure PostgreSQL Server Configurations.

https://learn.microsoft.com/en-us/rest/api/postgresql/singleserver/configurations/list-by-server?tabs=HTTP#configuration

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


