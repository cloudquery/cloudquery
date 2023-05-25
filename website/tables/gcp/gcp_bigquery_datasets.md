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
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id (PK)|utf8|
|access|json|
|creation_time|int64|
|dataset_reference|json|
|default_collation|utf8|
|default_encryption_configuration|json|
|default_partition_expiration_ms|int64|
|default_rounding_mode|utf8|
|default_table_expiration_ms|int64|
|description|utf8|
|etag|utf8|
|friendly_name|utf8|
|id (PK)|utf8|
|is_case_insensitive|bool|
|kind|utf8|
|labels|json|
|last_modified_time|int64|
|location|utf8|
|max_time_travel_hours|int64|
|satisfies_pzs|bool|
|self_link|utf8|
|storage_billing_model|utf8|
|tags|json|