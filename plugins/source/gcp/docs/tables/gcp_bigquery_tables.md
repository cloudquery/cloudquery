# Table: gcp_bigquery_tables



The primary key for this table is **_cq_id**.

## Relations
This table depends on [gcp_bigquery_datasets](gcp_bigquery_datasets.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|clone_definition|JSON|
|clustering|JSON|
|creation_time|Int|
|default_collation|String|
|description|String|
|encryption_configuration|JSON|
|etag|String|
|expiration_time|Int|
|external_data_configuration|JSON|
|friendly_name|String|
|id|String|
|kind|String|
|labels|JSON|
|last_modified_time|Int|
|location|String|
|materialized_view|JSON|
|max_staleness|String|
|model|JSON|
|num_bytes|Int|
|num_long_term_bytes|Int|
|num_physical_bytes|Int|
|num_rows|Int|
|num_active_logical_bytes|Int|
|num_active_physical_bytes|Int|
|num_long_term_logical_bytes|Int|
|num_long_term_physical_bytes|Int|
|num_partitions|Int|
|num_time_travel_physical_bytes|Int|
|num_total_logical_bytes|Int|
|num_total_physical_bytes|Int|
|range_partitioning|JSON|
|require_partition_filter|Bool|
|schema|JSON|
|self_link|String|
|snapshot_definition|JSON|
|streaming_buffer|JSON|
|table_reference|JSON|
|time_partitioning|JSON|
|type|String|
|view|JSON|