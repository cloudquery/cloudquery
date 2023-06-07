# Table: gcp_storage_buckets

This table shows data for GCP Storage Buckets.

https://cloud.google.com/storage/docs/json_api/v1/buckets#resource

The primary key for this table is **name**.

## Relations

The following tables depend on gcp_storage_buckets:
  - [gcp_storage_bucket_policies](gcp_storage_bucket_policies)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|name (PK)|`utf8`|
|acl|`json`|
|bucket_policy_only|`json`|
|uniform_bucket_level_access|`json`|
|public_access_prevention|`int64`|
|default_object_acl|`json`|
|default_event_based_hold|`bool`|
|predefined_acl|`utf8`|
|predefined_default_object_acl|`utf8`|
|location|`utf8`|
|custom_placement_config|`json`|
|meta_generation|`int64`|
|storage_class|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|versioning_enabled|`bool`|
|labels|`json`|
|requester_pays|`bool`|
|lifecycle|`json`|
|retention_policy|`json`|
|cors|`json`|
|encryption|`json`|
|logging|`json`|
|website|`json`|
|etag|`utf8`|
|location_type|`utf8`|
|project_number|`int64`|
|rpo|`int64`|
|autoclass|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that retention policies on log buckets are configured using Bucket Lock (Automated)

```sql
-- SELECT
-- gls.project_id,
-- gls.name AS "sink_name",
-- gsb.name AS "bucket_name",
-- gsb.retention_policy_is_locked,
-- gsb.retention_policy_retention_period,
-- gls.destination
-- FROM gcp_logging_sinks gls
-- JOIN gcp_storage_buckets gsb ON
-- gsb.name = REPLACE(gls.destination, 'storage.googleapis.com/', '')
-- WHERE gls.destination LIKE 'storage.googleapis.com/%'
-- AND ( gsb.retention_policy_is_locked = FALSE
-- OR gsb.retention_policy_retention_period = 0)

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

### Ensure that Cloud Storage buckets have uniform bucket-level access enabled (Automated)

```sql
-- SELECT project_id, name, self_link AS link
-- FROM gcp_storage_buckets
-- WHERE iam_configuration_uniform_bucket_level_access_enabled = FALSE;

SELECT
  name AS resource_id,
  'Ensure that Cloud Storage buckets have uniform bucket-level access enabled (Automated)'
    AS title,
  project_id AS project_id,
  CASE
  WHEN (uniform_bucket_level_access->>'Enabled')::BOOL = false THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_storage_buckets;
```


