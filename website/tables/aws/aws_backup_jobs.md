# Table: aws_backup_jobs

This table shows data for Backup Jobs.

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_BackupJob.html

The composite primary key for this table is (**account_id**, **region**, **backup_job_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|backup_job_id (PK)|`utf8`|
|backup_options|`json`|
|backup_size_in_bytes|`int64`|
|backup_type|`utf8`|
|backup_vault_arn|`utf8`|
|backup_vault_name|`utf8`|
|bytes_transferred|`int64`|
|completion_date|`timestamp[us, tz=UTC]`|
|created_by|`json`|
|creation_date|`timestamp[us, tz=UTC]`|
|expected_completion_date|`timestamp[us, tz=UTC]`|
|iam_role_arn|`utf8`|
|is_parent|`bool`|
|parent_job_id|`utf8`|
|percent_done|`utf8`|
|recovery_point_arn|`utf8`|
|resource_arn|`utf8`|
|resource_name|`utf8`|
|resource_type|`utf8`|
|start_by|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|status_message|`utf8`|