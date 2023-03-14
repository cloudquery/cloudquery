# Table: aws_lightsail_disk_snapshots

This table shows data for Lightsail Disk Snapshots.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_DiskSnapshot.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_lightsail_disks](aws_lightsail_disks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|disk_arn|String|
|tags|JSON|
|arn (PK)|String|
|created_at|Timestamp|
|from_disk_arn|String|
|from_disk_name|String|
|from_instance_arn|String|
|from_instance_name|String|
|is_from_auto_snapshot|Bool|
|location|JSON|
|name|String|
|progress|String|
|resource_type|String|
|size_in_gb|Int|
|state|String|
|support_code|String|