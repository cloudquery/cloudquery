# Table: aws_rds_engine_versions

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBEngineVersion.html

The composite primary key for this table is (**account_id**, **region**, **engine**, **engine_version**).

## Relations

The following tables depend on aws_rds_engine_versions:
  - [aws_rds_cluster_parameters](aws_rds_cluster_parameters.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|engine (PK)|String|
|engine_version (PK)|String|
|create_time|Timestamp|
|custom_db_engine_version_manifest|String|
|db_engine_description|String|
|db_engine_version_arn|String|
|db_engine_version_description|String|
|db_parameter_group_family|String|
|database_installation_files_s3_bucket_name|String|
|database_installation_files_s3_prefix|String|
|default_character_set|JSON|
|exportable_log_types|StringArray|
|kms_key_id|String|
|major_engine_version|String|
|status|String|
|supported_character_sets|JSON|
|supported_engine_modes|StringArray|
|supported_feature_names|StringArray|
|supported_nchar_character_sets|JSON|
|supported_timezones|JSON|
|supports_babelfish|Bool|
|supports_global_databases|Bool|
|supports_log_exports_to_cloudwatch_logs|Bool|
|supports_parallel_query|Bool|
|supports_read_replica|Bool|
|tag_list|JSON|
|valid_upgrade_target|JSON|