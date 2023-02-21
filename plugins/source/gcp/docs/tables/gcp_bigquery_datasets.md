# Table: gcp_bigquery_datasets

https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets#Dataset

The composite primary key for this table is (**project_id**, **id**).

## Relations

The following tables depend on gcp_bigquery_datasets:
  - [gcp_bigquery_tables](gcp_bigquery_tables.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|access|JSON|
|creation_time|Int|
|dataset_reference|JSON|
|default_collation|String|
|default_encryption_configuration|JSON|
|default_partition_expiration_ms|Int|
|default_table_expiration_ms|Int|
|description|String|
|etag|String|
|friendly_name|String|
|id (PK)|String|
|is_case_insensitive|Bool|
|kind|String|
|labels|JSON|
|last_modified_time|Int|
|location|String|
|max_time_travel_hours|Int|
|satisfies_pzs|Bool|
|self_link|String|
|storage_billing_model|String|
|tags|JSON|