
# Table: aws_fsx_data_repo_associations
The configuration of a data repository association that links an Amazon FSx for Lustre file system to an Amazon S3 bucket
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|association_id|text|The system-generated, unique ID of the data repository association|
|batch_import_meta_data_on_create|boolean|A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created|
|creation_time|timestamp without time zone|The time that the resource was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time|
|data_repository_path|text|The path to the Amazon S3 data repository that will be linked to the file system|
|failure_details_message|text|A detailed error message|
|file_system_id|text|The globally unique ID of the file system, assigned by Amazon FSx|
|file_system_path|text|A path on the file system that points to a high-level directory (such as /ns1/) or subdirectory (such as /ns1/subdir/) that will be mapped 1-1 with DataRepositoryPath|
|imported_file_chunk_size|bigint|For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk|
|lifecycle|text|Describes the state of a data repository association|
|arn|text|The Amazon Resource Name (ARN) for a given resource|
|s3_auto_export_policy_events|text[]|The AutoExportPolicy can have the following event values:  * NEW - Amazon FSx automatically exports new files and directories to the data repository as they are added to the file system  * CHANGED - Amazon FSx automatically exports changes to files and directories on the file system to the data repository  * DELETED - Files and directories are automatically deleted on the data repository when they are deleted on the file system  You can define any combination of event types for your AutoExportPolicy|
|s3_auto_import_policy_events|text[]|The AutoImportPolicy can have the following event values:  * NEW - Amazon FSx automatically imports metadata of files added to the linked S3 bucket that do not currently exist in the FSx file system  * CHANGED - Amazon FSx automatically updates file metadata and invalidates existing file content on the file system as files change in the data repository  * DELETED - Amazon FSx automatically deletes files on the file system as corresponding files are deleted in the data repository  You can define any combination of event types for your AutoImportPolicy|
|tags|jsonb|A list of Tag values, with a maximum of 50 elements|
