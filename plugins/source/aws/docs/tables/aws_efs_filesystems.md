# Table: aws_efs_filesystems

https://docs.aws.amazon.com/efs/latest/ug/API_FileSystemDescription.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|backup_policy_status|String|
|creation_time|Timestamp|
|creation_token|String|
|file_system_id|String|
|life_cycle_state|String|
|number_of_mount_targets|Int|
|owner_id|String|
|performance_mode|String|
|size_in_bytes|JSON|
|tags|JSON|
|availability_zone_id|String|
|availability_zone_name|String|
|encrypted|Bool|
|kms_key_id|String|
|name|String|
|provisioned_throughput_in_mibps|Float|
|throughput_mode|String|