# Table: aws_fsx_volumes

https://docs.aws.amazon.com/fsx/latest/APIReference/API_Volume.html

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
|administrative_actions|JSON|
|creation_time|Timestamp|
|file_system_id|String|
|lifecycle|String|
|lifecycle_transition_reason|JSON|
|name|String|
|ontap_configuration|JSON|
|open_zfs_configuration|JSON|
|tags|JSON|
|volume_id|String|
|volume_type|String|