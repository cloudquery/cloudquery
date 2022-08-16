
# Table: aws_fsx_filesystem_lustre_configuration
The configuration for the Amazon FSx for Lustre file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|filesystem_cq_id|uuid|Unique CloudQuery ID of aws_fsx_filesystems table (FK)|
|automatic_backup_retention_days|bigint|The number of days to retain automatic backups|
|copy_tags_to_backups|boolean|A boolean flag indicating whether tags on the file system are copied to backups. If it's set to true, all tags on the file system are copied to all automatic backups and any user-initiated backups where the user doesn't specify any tags. If this value is true, and you specify one or more tags, only the specified tags are copied to backups|
|daily_automatic_backup_start_time|text|A recurring daily time, in the format HH:MM|
|data_compression_type|text|The data compression configuration for the file system|
|data_repo_cfg_auto_import_policy|text|Describes the file system's linked S3 data repository's AutoImportPolicy|
|data_repo_cfg_export_path|text|The export path to the Amazon S3 bucket (and prefix) that you are using to store new and changed Lustre file system files in S3.|
|data_repo_cfg_failure_details_message|text|A detailed error message.|
|data_repo_cfg_import_path|text|The import path to the Amazon S3 bucket (and optional prefix) that you're using as the data repository for your FSx for Lustre file system, for example s3://import-bucket/optional-prefix|
|data_repo_cfg_imported_file_chunk_size|bigint|For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk|
|data_repo_cfg_lifecycle|text|Describes the state of the file system's S3 durable data repository, if it is configured with an S3 repository|
|deployment_type|text|The deployment type of the FSx for Lustre file system|
|drive_cache_type|text|The type of drive cache used by PERSISTENT_1 file systems that are provisioned with HDD storage devices|
|log_configuration_level|text|The data repository events that are logged by Amazon FSx.  * WARN_ONLY - only warning events are logged.  * ERROR_ONLY - only error events are logged.  * WARN_ERROR - both warning events and error events are logged.  * DISABLED - logging of data repository events is turned off.  This member is required.|
|log_configuration_destination|text|The Amazon Resource Name (ARN) that specifies the destination of the logs|
|mount_name|text|You use the MountName value when mounting the file system|
|per_unit_storage_throughput|bigint|Per unit storage throughput represents the megabytes per second of read or write throughput per 1 tebibyte of storage provisioned|
|root_squash_configuration_no_squash_nids|text[]|When root squash is enabled, you can optionally specify an array of NIDs of clients for which root squash does not apply|
|root_squash_configuration_root_squash|text|You enable root squash by setting a user ID (UID) and group ID (GID) for the file system in the format UID:GID (for example, 365534:65534)|
|weekly_maintenance_start_time|text|The preferred start time to perform weekly maintenance, formatted d:HH:MM in the UTC time zone|
