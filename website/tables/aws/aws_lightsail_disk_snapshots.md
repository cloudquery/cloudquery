# Table: aws_lightsail_disk_snapshots

This table shows data for Lightsail Disk Snapshots.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_DiskSnapshot.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_lightsail_disks](aws_lightsail_disks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|disk_arn|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|from_disk_arn|`utf8`|
|from_disk_name|`utf8`|
|from_instance_arn|`utf8`|
|from_instance_name|`utf8`|
|is_from_auto_snapshot|`bool`|
|location|`json`|
|name|`utf8`|
|progress|`utf8`|
|resource_type|`utf8`|
|size_in_gb|`int64`|
|state|`utf8`|
|support_code|`utf8`|