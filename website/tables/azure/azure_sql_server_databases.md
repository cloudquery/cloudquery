# Table: azure_sql_server_databases

This table shows data for Azure SQL Server Databases.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/databases/list-by-server?tabs=HTTP#database

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_servers](azure_sql_servers).

The following tables depend on azure_sql_server_databases:
  - [azure_sql_server_database_blob_auditing_policies](azure_sql_server_database_blob_auditing_policies)
  - [azure_sql_server_database_long_term_retention_policies](azure_sql_server_database_long_term_retention_policies)
  - [azure_sql_server_database_threat_protections](azure_sql_server_database_threat_protections)
  - [azure_sql_server_database_vulnerability_assessments](azure_sql_server_database_vulnerability_assessments)
  - [azure_sql_transparent_data_encryptions](azure_sql_transparent_data_encryptions)

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
|kind|`utf8`|
|managed_by|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that "Auditing" is set to "On" (Automated)

```sql
SELECT
  'Ensure that "Auditing" is set to "On" (Automated)' AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN assdbap.properties->>'state' != 'Enabled' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_databases AS assd ON s._cq_id = assd._cq_parent_id
  LEFT JOIN azure_sql_server_database_blob_auditing_policies AS assdbap ON
      assd._cq_id = assdbap._cq_parent_id;
```

### Ensure that "Data encryption" is set to "On" on a SQL Database (Automated)

```sql
SELECT
  'Ensure that "Data encryption" is set to "On" on a SQL Database (Automated)'
    AS title,
  s.subscription_id,
  asd.id AS database_id,
  CASE
  WHEN tde.properties->>'state' IS DISTINCT FROM 'Enabled' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_databases AS asd ON s._cq_id = asd._cq_parent_id
  LEFT JOIN azure_sql_transparent_data_encryptions AS tde ON
      asd._cq_id = tde._cq_parent_id
WHERE
  asd.name != 'master';
```

### Long-term geo-redundant backup should be enabled for Azure SQL Databases

```sql
SELECT
  'Long-term geo-redundant backup should be enabled for Azure SQL Databases'
    AS title,
  s.subscription_id,
  rp.id,
  CASE
  WHEN rp.id IS NULL
  OR (
      rp.properties->>'weeklyRetention' IS NOT DISTINCT FROM 'PT0S'
      AND rp.properties->>'monthlyRetention' IS NOT DISTINCT FROM 'PT0S'
      AND rp.properties->>'yearlyRetention' IS NOT DISTINCT FROM 'PT0S'
    )
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_databases AS asd ON s._cq_id = asd._cq_parent_id
  LEFT JOIN azure_sql_server_database_long_term_retention_policies AS rp ON
      asd._cq_id = rp._cq_parent_id;
```

### SQL databases should have vulnerability findings resolved

```sql
WITH
  safe_dbs
    AS (
      SELECT
        s.id AS sql_database_id
      FROM
        azure_sql_server_database_vulnerability_assessment_scans AS s
        JOIN (
            SELECT
              _cq_id, max((properties->>'endTime')::TIMESTAMP) AS max_dt
            FROM
              azure_sql_server_database_vulnerability_assessment_scans
            GROUP BY
              _cq_id
          )
            AS t ON
            s._cq_id = t._cq_id
            AND (properties->>'endTime')::TIMESTAMP = t.max_dt
      WHERE
        (s.properties->>'numberOfFailedSecurityChecks')::INT8 = 0
    )
SELECT
  'SQL databases should have vulnerability findings resolved' AS title,
  s.subscription_id,
  d.id,
  CASE
  WHEN d.id IS NULL THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_databases AS d ON s._cq_id = d._cq_parent_id
  LEFT JOIN safe_dbs AS sd ON d.id = sd.sql_database_id;
```


