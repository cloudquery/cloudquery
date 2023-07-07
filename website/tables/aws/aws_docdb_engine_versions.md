# Table: aws_docdb_engine_versions

This table shows data for Amazon DocumentDB Engine Versions.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBEngineVersion.html

The composite primary key for this table is (**account_id**, **region**, **engine**, **engine_version**).

## Relations

The following tables depend on aws_docdb_engine_versions:
  - [aws_docdb_cluster_parameters](aws_docdb_cluster_parameters)
  - [aws_docdb_orderable_db_instance_options](aws_docdb_orderable_db_instance_options)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|db_engine_description|`utf8`|
|db_engine_version_description|`utf8`|
|db_parameter_group_family|`utf8`|
|engine (PK)|`utf8`|
|engine_version (PK)|`utf8`|
|exportable_log_types|`list<item: utf8, nullable>`|
|supports_log_exports_to_cloudwatch_logs|`bool`|
|valid_upgrade_target|`json`|