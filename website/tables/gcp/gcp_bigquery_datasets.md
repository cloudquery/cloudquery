# Table: gcp_bigquery_datasets

This table shows data for GCP BigQuery Datasets.

https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets#Dataset

The composite primary key for this table is (**project_id**, **id**).

## Relations

The following tables depend on gcp_bigquery_datasets:
  - [gcp_bigquery_tables](gcp_bigquery_tables)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|access|`json`|
|creation_time|`int64`|
|dataset_reference|`json`|
|default_collation|`utf8`|
|default_encryption_configuration|`json`|
|default_partition_expiration_ms|`int64`|
|default_rounding_mode|`utf8`|
|default_table_expiration_ms|`int64`|
|description|`utf8`|
|etag|`utf8`|
|friendly_name|`utf8`|
|id (PK)|`utf8`|
|is_case_insensitive|`bool`|
|kind|`utf8`|
|labels|`json`|
|last_modified_time|`int64`|
|location|`utf8`|
|max_time_travel_hours|`int64`|
|satisfies_pzs|`bool`|
|self_link|`utf8`|
|storage_billing_model|`utf8`|
|tags|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that BigQuery datasets are not anonymously or publicly accessible (Automated)

```sql
-- SELECT
-- d.project_id,
-- d.id,
-- d.friendly_name,
-- d.self_link AS dataset_link,
-- a.special_group AS "group",
-- a."role"
-- FROM gcp_bigquery_datasets d
-- JOIN gcp_bigquery_dataset_accesses a ON
-- d.id = a.dataset_id
-- WHERE a."role" = 'allUsers'
-- OR a."role" = 'allAuthenticatedUsers';

SELECT
  DISTINCT
  d.id AS resource_id,
  'Ensure that BigQuery datasets are not anonymously or publicly accessible (Automated)'
    AS title,
  d.project_id AS project_id,
  CASE
  WHEN a->>'role' = 'allUsers' OR a->>'role' = 'allAuthenticatedUsers'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_bigquery_datasets AS d, jsonb_array_elements(d.access) AS a;
```

### Ensure that all BigQuery Tables are encrypted with Customer-managed encryption key (CMEK) (Automated)

```sql
-- SELECT project_id, id, friendly_name, self_link AS link
-- FROM gcp_bigquery_datasets
-- WHERE default_encryption_configuration_kms_key_name = ''
-- OR default_encryption_configuration_kms_key_name IS NULL;

SELECT
  d.id AS resource_id,
  'Ensure that all BigQuery Tables are encrypted with Customer-managed encryption key (CMEK) (Automated)'
    AS title,
  d.project_id AS project_id,
  CASE
  WHEN d.default_encryption_configuration->>'kmsKeyName' = ''
  OR (d.default_encryption_configuration->>'kmsKeyName') IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_bigquery_datasets AS d;
```

### Ensure that a Default Customer-managed encryption key (CMEK) is specified for all BigQuery Data Sets (Automated)

```sql
-- SELECT d.project_id, d.id, d.friendly_name, d.self_link AS dataset_link, t.self_link AS table_link
-- FROM gcp_bigquery_datasets d
-- JOIN gcp_bigquery_dataset_tables t ON
-- d.id = t.dataset_id
-- WHERE encryption_configuration_kms_key_name = '' OR default_encryption_configuration_kms_key_name IS NULL;

SELECT
  DISTINCT
  d.id AS resource_id,
  'Ensure that a Default Customer-managed encryption key (CMEK) is specified for all BigQuery Data Sets (Automated)'
    AS title,
  d.project_id AS project_id,
  CASE
  WHEN t.encryption_configuration->>'kmsKeyName' = ''
  OR (d.default_encryption_configuration->>'kmsKeyName') IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_bigquery_datasets AS d
  JOIN gcp_bigquery_tables AS t ON
      d.dataset_reference->>'datasetId' = t.table_reference->>'datasetId'
      AND d.dataset_reference->>'projectId' = t.table_reference->>'projectId';
```


