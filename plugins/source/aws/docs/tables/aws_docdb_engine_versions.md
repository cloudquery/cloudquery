# Table: aws_docdb_engine_versions

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBEngineVersion.html

The composite primary key for this table is (**account_id**, **engine**, **engine_version**).

## Relations

The following tables depend on aws_docdb_engine_versions:
  - [aws_docdb_cluster_parameters](aws_docdb_cluster_parameters.md)
  - [aws_docdb_orderable_db_instance_options](aws_docdb_orderable_db_instance_options.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|engine (PK)|String|
|engine_version (PK)|String|
|db_engine_description|String|
|db_engine_version_description|String|
|db_parameter_group_family|String|
|exportable_log_types|StringArray|
|supports_log_exports_to_cloudwatch_logs|Bool|
|valid_upgrade_target|JSON|