# Table: gcp_bigquery_tables

This table shows data for GCP BigQuery Tables.

https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#Table

The composite primary key for this table is (**project_id**, **id**).

## Relations

This table depends on [gcp_bigquery_datasets](gcp_bigquery_datasets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|clone_definition|`json`|
|clustering|`json`|
|creation_time|`int64`|
|default_collation|`utf8`|
|default_rounding_mode|`utf8`|
|description|`utf8`|
|encryption_configuration|`json`|
|etag|`utf8`|
|expiration_time|`int64`|
|external_data_configuration|`json`|
|friendly_name|`utf8`|
|id (PK)|`utf8`|
|kind|`utf8`|
|labels|`json`|
|last_modified_time|`int64`|
|location|`utf8`|
|materialized_view|`json`|
|max_staleness|`utf8`|
|model|`json`|
|num_bytes|`int64`|
|num_long_term_bytes|`int64`|
|num_physical_bytes|`int64`|
|num_rows|`int64`|
|num_active_logical_bytes|`int64`|
|num_active_physical_bytes|`int64`|
|num_long_term_logical_bytes|`int64`|
|num_long_term_physical_bytes|`int64`|
|num_partitions|`int64`|
|num_time_travel_physical_bytes|`int64`|
|num_total_logical_bytes|`int64`|
|num_total_physical_bytes|`int64`|
|range_partitioning|`json`|
|require_partition_filter|`bool`|
|schema|`json`|
|self_link|`utf8`|
|snapshot_definition|`json`|
|streaming_buffer|`json`|
|table_reference|`json`|
|time_partitioning|`json`|
|type|`utf8`|
|view|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that a Default Customer-managed encryption key (CMEK) is specified for all BigQuery Data Sets (Automated)

```sql
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


