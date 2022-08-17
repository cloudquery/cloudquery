
# Table: aws_fsx_filesystem_open_zfs_configuration
The configuration for the Amazon FSx for OpenZFS file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|filesystem_cq_id|uuid|Unique CloudQuery ID of aws_fsx_filesystems table (FK)|
|automatic_backup_retention_days|bigint|The number of days to retain automatic backups|
|copy_tags_to_backups|boolean|A Boolean value indicating whether tags on the file system should be copied to backups|
|copy_tags_to_volumes|boolean|A Boolean value indicating whether tags for the volume should be copied to snapshots|
|daily_automatic_backup_start_time|text|A recurring daily time, in the format HH:MM|
|deployment_type|text|Specifies the file-system deployment type|
|disk_iops_configuration_iops|bigint|The total number of SSD IOPS provisioned for the file system.|
|disk_iops_configuration_mode|text|Specifies whether the number of IOPS for the file system is using the system default (AUTOMATIC) or was provisioned by the customer (USER_PROVISIONED).|
|root_volume_id|text|The ID of the root volume of the OpenZFS file system.|
|throughput_capacity|bigint|The throughput of an Amazon FSx file system, measured in megabytes per second (MBps)|
|weekly_maintenance_start_time|text|A recurring weekly time, in the format D:HH:MM|
