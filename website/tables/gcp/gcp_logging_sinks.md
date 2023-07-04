# Table: gcp_logging_sinks

This table shows data for GCP Logging Sinks.

https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.sinks#LogSink

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|destination|`utf8`|
|filter|`utf8`|
|description|`utf8`|
|disabled|`bool`|
|exclusions|`json`|
|output_version_format|`utf8`|
|writer_identity|`utf8`|
|include_children|`bool`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that retention policies on log buckets are configured using Bucket Lock (Automated)

```sql
SELECT
  DISTINCT
  gsb.name AS resource_id,
  'Ensure that retention policies on log buckets are configured using Bucket Lock (Automated)'
    AS title,
  gls.project_id AS project_id,
  CASE
  WHEN gls.destination LIKE 'storage.googleapis.com/%'
  AND (
      (gsb.retention_policy->>'IsLocked')::BOOL = false
      OR (gsb.retention_policy->>'RetentionPeriod')::INT8 = 0
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_logging_sinks AS gls
  JOIN gcp_storage_buckets AS gsb ON
      gsb.name = replace(gls.destination, 'storage.googleapis.com/', '');
```

### Ensure that sinks are configured for all log entries (Automated)

```sql
WITH
  found_sinks
    AS (
      SELECT
        project_id, name, count(*) AS configured_sinks
      FROM
        gcp_logging_sinks AS gls
      WHERE
        gls.filter = ''
      GROUP BY
        project_id, name
    )
SELECT
  name AS resource_id,
  'Ensure that sinks are configured for all log entries (Automated)' AS title,
  project_id AS project_id,
  CASE WHEN configured_sinks = 0 THEN 'fail' ELSE 'pass' END AS status
FROM
  found_sinks;
```


