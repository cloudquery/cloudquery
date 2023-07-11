# Table: aws_rds_engine_versions

This table shows data for Amazon Relational Database Service (RDS) Engine Versions.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html

The composite primary key for this table is (**account_id**, **region**, **_engine_version_hash**).

## Relations

The following tables depend on aws_rds_engine_versions:
  - [aws_rds_cluster_parameters](aws_rds_cluster_parameters)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|_engine_version_hash (PK)|`utf8`|
|tag_list|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|custom_db_engine_version_manifest|`utf8`|
|db_engine_description|`utf8`|
|db_engine_media_type|`utf8`|
|db_engine_version_arn|`utf8`|
|db_engine_version_description|`utf8`|
|db_parameter_group_family|`utf8`|
|database_installation_files_s3_bucket_name|`utf8`|
|database_installation_files_s3_prefix|`utf8`|
|default_character_set|`json`|
|engine|`utf8`|
|engine_version|`utf8`|
|exportable_log_types|`list<item: utf8, nullable>`|
|image|`json`|
|kms_key_id|`utf8`|
|major_engine_version|`utf8`|
|status|`utf8`|
|supported_ca_certificate_identifiers|`list<item: utf8, nullable>`|
|supported_character_sets|`json`|
|supported_engine_modes|`list<item: utf8, nullable>`|
|supported_feature_names|`list<item: utf8, nullable>`|
|supported_nchar_character_sets|`json`|
|supported_timezones|`json`|
|supports_babelfish|`bool`|
|supports_certificate_rotation_without_restart|`bool`|
|supports_global_databases|`bool`|
|supports_log_exports_to_cloudwatch_logs|`bool`|
|supports_parallel_query|`bool`|
|supports_read_replica|`bool`|
|valid_upgrade_target|`json`|