# Table: aws_docdb_engine_versions

This table shows data for Amazon DocumentDB Engine Versions.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBEngineVersion.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **engine**, **engine_version**).
## Relations

The following tables depend on aws_docdb_engine_versions:
  - [aws_docdb_cluster_parameters](aws_docdb_cluster_parameters.md)
  - [aws_docdb_orderable_db_instance_options](aws_docdb_orderable_db_instance_options.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|db_engine_description|`utf8`|
|db_engine_version_description|`utf8`|
|db_parameter_group_family|`utf8`|
|engine|`utf8`|
|engine_version|`utf8`|
|exportable_log_types|`list<item: utf8, nullable>`|
|supported_ca_certificate_identifiers|`list<item: utf8, nullable>`|
|supports_certificate_rotation_without_restart|`bool`|
|supports_log_exports_to_cloudwatch_logs|`bool`|
|valid_upgrade_target|`json`|