
# Table: gcp_bigquery_dataset_tables
Model options used for the first training run These options are immutable for subsequent training runs Default values are used for any options not specified in the input query
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|dataset_id|uuid||
|clustering_fields|text[]|One or more fields on which data should be clustered Only top-level, non-repeated, simple-type fields are supported When you cluster a table using multiple columns, the order of columns you specify is important The order of the specified columns determines the sort order of the data|
|creation_time|bigint|The time when this table was created, in milliseconds since the epoch|
|description|text|A user-friendly description of this table|
|encryption_configuration_kms_key_name|text|Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table The BigQuery Service Account associated with your project requires access to this encryption key|
|etag|text|A hash of the table metadata Used to ensure there were no concurrent modifications to the resource when attempting an update Not guaranteed to change when the table contents or the fields numRows, numBytes, numLongTermBytes or lastModifiedTime change|
|expiration_time|bigint|The time when this table expires, in milliseconds since the epoch If not present, the table will persist indefinitely Expired tables will be deleted and their storage reclaimed The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables|
|external_data_configuration_autodetect|boolean|Try to detect schema and format options automatically Any option specified explicitly will be honored|
|external_data_configuration_compression|text|The compression type of the data source Possible values include GZIP and NONE The default value is NONE This setting is ignored for Google Cloud Bigtable, Google Cloud Datastore backups and Avro formats|
|external_data_configuration_connection_id|text|Connection for external data source|
|external_data_configuration_ignore_unknown_values|boolean|Indicates if BigQuery should allow extra values that are not represented in the table schema|
|external_data_configuration_max_bad_records|bigint|The maximum number of bad records that BigQuery can ignore when reading data If the number of bad records exceeds this value, an invalid error is returned in the job result This is only valid for CSV, JSON, and Google Sheets The default value is 0, which requires that all records are valid This setting is ignored for Google Cloud Bigtable, Google Cloud Datastore backups and Avro formats|
|external_data_configuration_schema|jsonb|The schema for the data Schema is required for CSV and JSON formats Schema is disallowed for Google Cloud Bigtable, Cloud Datastore backups, and Avro formats|
|external_data_configuration_source_format|text|The data format For CSV files, specify "CSV" For Google sheets, specify "GOOGLE_SHEETS" For newline-delimited JSON, specify "NEWLINE_DELIMITED_JSON" For Avro files, specify "AVRO" For Google Cloud Datastore backups, specify "DATASTORE_BACKUP" [Beta] For Google Cloud Bigtable, specify "BIGTABLE"|
|external_data_configuration_source_uris|text[]|The fully-qualified URIs that point to your data in Google Cloud For Google Cloud Storage URIs: Each URI can contain one '*' wildcard character and it must come after the 'bucket' name Size limits related to load jobs apply to external data sources For Google Cloud Bigtable URIs: Exactly one URI can be specified and it has be a fully specified and valid HTTPS URL for a Google Cloud Bigtable table For Google Cloud Datastore backups, exactly one URI can be specified Also, the '*' wildcard character is not allowed|
|friendly_name|text|A descriptive name for this table|
|table_id|text|An opaque ID uniquely identifying the table|
|kind|text|The type of the resource|
|labels|jsonb|The labels associated with this table You can use these to organize and group your tables Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes International characters are allowed Label values are optional Label keys must start with a letter and each label in the list must have a different key|
|last_modified_time|bigint|The time when this table was last modified, in milliseconds since the epoch|
|location|text|The geographic location where the table resides This value is inherited from the dataset|
|materialized_view_enable_refresh|boolean|Enable automatic refresh of the materialized view when the base table is updated The default value is "true"|
|materialized_view_last_refresh_time|bigint|The time when this materialized view was last modified, in milliseconds since the epoch|
|materialized_view_query|text|A query whose result is persisted|
|materialized_view_refresh_interval_ms|bigint|The maximum frequency at which this materialized view will be refreshed The default value is "1800000" (30 minutes)|
|model_options_labels|text[]||
|model_options_loss_type|text||
|model_options_model_type|text||
|num_bytes|bigint|The size of this table in bytes, excluding any data in the streaming buffer|
|num_long_term_bytes|bigint|The number of bytes in the table that are considered "long-term storage"|
|num_physical_bytes|bigint|The physical size of this table in bytes, excluding any data in the streaming buffer This includes compression and storage used for time travel|
|num_rows|bigint|The number of rows of data in this table, excluding any data in the streaming buffer|
|range_partitioning_field|text|The table is partitioned by this field The field must be a top-level NULLABLE/REQUIRED field The only supported type is INTEGER/INT64|
|range_partitioning_range_end|bigint|The end of range partitioning, exclusive|
|range_partitioning_range_interval|bigint|The width of each interval|
|range_partitioning_range_start|bigint|The start of range partitioning, inclusive|
|require_partition_filter|boolean|If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified|
|schema|jsonb|Describes the schema of this table|
|self_link|text|A URL that can be used to access this resource again|
|streaming_buffer_estimated_bytes|bigint|A lower-bound estimate of the number of bytes currently in the streaming buffer|
|streaming_buffer_estimated_rows|bigint|A lower-bound estimate of the number of rows currently in the streaming buffer|
|streaming_buffer_oldest_entry_time|bigint|Contains the timestamp of the oldest entry in the streaming buffer, in milliseconds since the epoch, if the streaming buffer is available|
|time_partitioning_expiration_ms|bigint|Number of milliseconds for which to keep the storage for partitions in the table The storage in a partition will have an expiration time of its partition time plus this value|
|time_partitioning_field|text|If not set, the table is partitioned by pseudo column, referenced via either '_PARTITIONTIME' as TIMESTAMP type, or '_PARTITIONDATE' as DATE type If field is specified, the table is instead partitioned by this field The field must be a top-level TIMESTAMP or DATE field Its mode must be NULLABLE or REQUIRED|
|time_partitioning_require_partition_filter|boolean||
|time_partitioning_type|text|The supported types are DAY, HOUR, MONTH, and YEAR, which will generate one partition per day, hour, month, and year, respectively When the type is not specified, the default behavior is DAY|
|type|text|Describes the table type The following values are supported: TABLE: A normal BigQuery table VIEW: A virtual table defined by a SQL query SNAPSHOT: An immutable, read-only table that is a copy of another table MATERIALIZED_VIEW: SQL query whose result is persisted EXTERNAL: A table that references data stored in an external storage system, such as Google Cloud Storage The default value is TABLE|
|view_query|text|A query that BigQuery executes when the view is referenced|
|view_use_legacy_sql|boolean|Specifies whether to use BigQuery's legacy SQL for this view The default value is true If set to false, the view will use BigQuery's standard SQL: https://cloudgooglecom/bigquery/sql-reference/ Queries and views that reference this view must use the same flag value|
