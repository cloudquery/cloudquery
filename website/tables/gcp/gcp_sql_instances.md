# Table: gcp_sql_instances

This table shows data for GCP SQL Instances.

https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/instances#DatabaseInstance

The primary key for this table is **self_link**.

## Relations

The following tables depend on gcp_sql_instances:
  - [gcp_sql_users](gcp_sql_users)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|available_maintenance_versions|`list<item: utf8, nullable>`|
|backend_type|`utf8`|
|connection_name|`utf8`|
|create_time|`utf8`|
|current_disk_size|`int64`|
|database_installed_version|`utf8`|
|database_version|`utf8`|
|disk_encryption_configuration|`json`|
|disk_encryption_status|`json`|
|etag|`utf8`|
|failover_replica|`json`|
|gce_zone|`utf8`|
|instance_type|`utf8`|
|ip_addresses|`json`|
|ipv6_address|`utf8`|
|kind|`utf8`|
|maintenance_version|`utf8`|
|master_instance_name|`utf8`|
|max_disk_size|`int64`|
|name|`utf8`|
|on_premises_configuration|`json`|
|out_of_disk_report|`json`|
|project|`utf8`|
|region|`utf8`|
|replica_configuration|`json`|
|replica_names|`list<item: utf8, nullable>`|
|root_password|`utf8`|
|satisfies_pzs|`bool`|
|scheduled_maintenance|`json`|
|secondary_gce_zone|`utf8`|
|self_link (PK)|`utf8`|
|server_ca_cert|`json`|
|service_account_email_address|`utf8`|
|settings|`json`|
|state|`utf8`|
|suspension_reason|`list<item: utf8, nullable>`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that Cloud SQL database instances are not open to the world (Automated)

```sql
SELECT
  DISTINCT
  gsi.name AS resource_id,
  'Ensure that Cloud SQL database instances are not open to the world (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'SQLSERVER%'
  AND gsisican->>'value' = '0.0.0.0/0'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi,
  jsonb_array_elements(gsi.settings->'ipConfiguration'->'authorizedNetworks')
    AS gsisican;
```

### Ensure that Cloud SQL database instances do not have public IPs (Automated)

```sql
SELECT
  DISTINCT
  gsi.name AS resource_id,
  'Ensure that Cloud SQL database instances do not have public IPs (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'SQLSERVER%'
  AND gsiia->>'type' = 'PRIMARY'
  OR gsi.backend_type != 'SECOND_GEN'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi, jsonb_array_elements(gsi.ip_addresses) AS gsiia;
```

### Ensure that the Cloud SQL database instance requires all incoming connections to use SSL (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'SQLSERVER%'
-- AND settings_ip_configuration_require_ssl = FALSE;

SELECT
  gsi.name AS resource_id,
  'Ensure that the Cloud SQL database instance requires all incoming connections to use SSL (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'SQLSERVER%'
  AND (gsi.settings->'ipConfiguration'->>'requireSsl')::BOOL = false
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi;
```

### Ensure that Cloud SQL database instances are configured with automated backups (Automated)

```sql
-- SELECT project_id, name, self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'SQLSERVER%'
-- AND settings_backup_enabled = FALSE;

SELECT
  gsi.name AS resource_id,
  'Ensure that Cloud SQL database instances are configured with automated backups (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'SQLSERVER%'
  AND (gsi.settings->'backupConfiguration'->>'enabled')::BOOL = false
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi;
```

### Ensure that the "local_infile" database flag for a Cloud SQL Mysql instance is set to "off" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'MYSQL%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'local_infile' != 'off'
-- OR settings_database_flags ->> 'local_infile' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure that the "local_infile" database flag for a Cloud SQL Mysql instance is set to "off" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'MYSQL%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'on')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'local_infile';
```

### Ensure "skip_show_database" database flag for Cloud SQL Mysql instance is set to "on" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'MYSQL%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'skip_show_database' != 'on'
-- OR settings_database_flags ->> 'skip_show_database' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure "skip_show_database" database flag for Cloud SQL Mysql instance is set to "on" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'MYSQL%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'on')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'skip_show_database';
```

### Ensure that the "log_checkpoints" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_checkpoints' != 'on'
-- OR settings_database_flags ->> 'log_checkpoints' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure that the "log_checkpoints" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'on')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_checkpoints';
```

### Ensure that the "log_connections" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_connections' != 'on'
-- OR settings_database_flags ->> 'log_connections' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure that the "log_connections" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'on')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_connections';
```

### Ensure that the log_disconnections" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_disconnections' != 'on'
-- OR settings_database_flags ->> 'log_disconnections' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure that the log_disconnections" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'on')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_disconnections';
```

### Ensure "log_duration" database flag for Cloud SQL PostgreSQL instance is set to "on" (Manual)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_duration' != 'on'
-- OR settings_database_flags ->> 'log_duration' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure "log_duration" database flag for Cloud SQL PostgreSQL instance is set to "on" (Manual)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'on')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_duration';
```

### Ensure "log_error_verbosity" database flag for Cloud SQL PostgreSQL instance is set to "DEFAULT" or stricter (Manual)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_error_verbosity' NOT IN('default', 'terse')
-- OR settings_database_flags ->> 'log_error_verbosity' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure "log_error_verbosity" database flag for Cloud SQL PostgreSQL instance is set to "DEFAULT" or stricter (Manual)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' NOT IN ('default', 'terse'))
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_error_verbosity';
```

### Ensure "log_executor_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_executor_stats' != 'off'
-- OR settings_database_flags ->> 'log_executor_stats' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure "log_executor_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'off')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_executor_stats';
```

### Ensure "log_hostname" database flag for Cloud SQL PostgreSQL instance is set appropriately (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_hostname' != 'on'
-- OR settings_database_flags ->> 'log_hostname' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure "log_hostname" database flag for Cloud SQL PostgreSQL instance is set appropriately (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'on')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_hostname';
```

### Ensure that the "log_lock_waits" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags IS NULL OR settings_database_flags ->> 'log_lock_waits' != 'on'
-- OR settings_database_flags ->> 'log_lock_waits' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure that the "log_lock_waits" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'on')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_lock_waits';
```

### Ensure that the "log_min_duration_statement" database flag for Cloud SQL PostgreSQL instance is set to "-1" (disabled) (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_min_duration_statement' != '-1'
-- OR settings_database_flags ->> 'log_min_duration_statement' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure that the "log_min_duration_statement" database flag for Cloud SQL PostgreSQL instance is set to "-1" (disabled) (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' != '-1')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_min_duration_statement';
```

### Ensure that the "log_min_messages" database flag for Cloud SQL PostgreSQL instance is set appropriately (Manual)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_min_error_statement' NOT IN('error', 'log', 'fatal', 'panic')
-- OR settings_database_flags ->> 'log_min_error_statement' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure that the "log_min_messages" database flag for Cloud SQL PostgreSQL instance is set appropriately (Manual)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND (
      (f->>'value') IS NULL
      OR f->>'value' NOT IN ('error', 'log', 'fatal', 'panic')
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_min_error_statement';
```

### Ensure "log_parser_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_parser_stats' != 'off'
-- OR settings_database_flags ->> 'log_parser_stats' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure "log_parser_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'off')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_parser_stats';
```

### Ensure "log_planner_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_planner_stats' != 'off'
-- OR settings_database_flags ->> 'log_planner_stats' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure "log_planner_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'off')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_planner_stats';
```

### Ensure "log_statement_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_statement_stats' != 'off'
-- OR settings_database_flags ->> 'log_statement_stats' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure "log_statement_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'off')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_statement_stats';
```

### Ensure that the "log_temp_files" database flag for Cloud SQL PostgreSQL instance is set to "0" (on) (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'POSTGRES%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'log_temp_files' != '0'
-- OR settings_database_flags ->> 'log_temp_files' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure that the "log_temp_files" database flag for Cloud SQL PostgreSQL instance is set to "0" (on) (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'POSTGRES%'
  AND ((f->>'value') IS NULL OR f->>'value' != '0')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'log_temp_files';
```

### Ensure that the "contained database authentication" database flag for Cloud SQL on the SQL Server instance is set to "off" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'SQLSERVER%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'contained database authentication' != 'off'
-- OR settings_database_flags ->> 'contained database authentication' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure that the "contained database authentication" database flag for Cloud SQL on the SQL Server instance is set to "off" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'SQLSERVER%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'off')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'contained database authentication';
```

### Ensure that the "cross db ownership chaining" database flag for Cloud SQL SQL Server instance is set to "off" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'SQLSERVER%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'cross db ownership chaining' != 'off'
-- OR settings_database_flags ->> 'cross db ownership chaining' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure that the "cross db ownership chaining" database flag for Cloud SQL SQL Server instance is set to "off" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'SQLSERVER%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'off')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'cross db ownership chaining';
```

### Ensure "external scripts enabled" database flag for Cloud SQL SQL Server instance is set to "off" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'SQLSERVER%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'external scripts enabled' != 'off'
-- OR settings_database_flags ->> 'external scripts enabled' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure "external scripts enabled" database flag for Cloud SQL SQL Server instance is set to "off" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'SQLSERVER%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'off')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'external scripts enabled';
```

### Ensure "remote access" database flag for Cloud SQL SQL Server instance is set to "off" (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'SQLSERVER%'
-- AND (settings_database_flags IS NULL
-- OR settings_database_flags ->> 'remote access' != 'off'
-- OR settings_database_flags ->> 'remote access' IS NULL);

SELECT
  gsi.name AS resource_id,
  'Ensure "remote access" database flag for Cloud SQL SQL Server instance is set to "off" (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'SQLSERVER%'
  AND ((f->>'value') IS NULL OR f->>'value' != 'off')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'remote access';
```

### Ensure "user connections" database flag for Cloud SQL SQL Server instance is set as appropriate (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'SQLSERVER%'
-- AND settings_database_flags IS NULL
-- OR settings_database_flags ->> 'user connections' IS NULL;

SELECT
  gsi.name AS resource_id,
  'Ensure "user connections" database flag for Cloud SQL SQL Server instance is set as appropriate (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'SQLSERVER%' AND (f->>'value') IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'user connections';
```

### Ensure "user options" database flag for Cloud SQL SQL Server instance is not configured (Automated)

```sql
-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'SQLSERVER%'
-- AND settings_database_flags IS NULL
-- OR settings_database_flags ->> 'user options' IS NOT NULL;

SELECT
  gsi.name AS resource_id,
  'Ensure "user options" database flag for Cloud SQL SQL Server instance is not configured (Automated)'
    AS title,
  gsi.project_id AS project_id,
  CASE
  WHEN gsi.database_version LIKE 'SQLSERVER%' AND (f->>'value') IS NOT NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_sql_instances AS gsi
  LEFT JOIN jsonb_array_elements(gsi.settings->'databaseFlags') AS f ON
      f->>'name' = 'user options';
```


