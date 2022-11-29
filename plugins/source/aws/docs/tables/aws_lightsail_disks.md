# Table: aws_lightsail_disks

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Disk.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_lightsail_disks:
  - [aws_lightsail_disk_snapshots](aws_lightsail_disk_snapshots.md)

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
|tags|JSON|
|add_ons|JSON|
|attached_to|String|
|attachment_state|String|
|created_at|Timestamp|
|gb_in_use|Int|
|iops|Int|
|is_attached|Bool|
|is_system_disk|Bool|
|location|JSON|
|name|String|
|path|String|
|resource_type|String|
|size_in_gb|Int|
|state|String|
|support_code|String|