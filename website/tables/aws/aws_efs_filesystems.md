# Table: aws_efs_filesystems

This table shows data for Amazon Elastic File System (EFS) Filesystems.

https://docs.aws.amazon.com/efs/latest/ug/API_FileSystemDescription.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|backup_policy_status|utf8|
|tags|json|
|creation_time|timestamp[us, tz=UTC]|
|creation_token|utf8|
|file_system_id|utf8|
|life_cycle_state|utf8|
|number_of_mount_targets|int64|
|owner_id|utf8|
|performance_mode|utf8|
|size_in_bytes|json|
|availability_zone_id|utf8|
|availability_zone_name|utf8|
|encrypted|bool|
|file_system_arn|utf8|
|kms_key_id|utf8|
|name|utf8|
|provisioned_throughput_in_mibps|float64|
|throughput_mode|utf8|