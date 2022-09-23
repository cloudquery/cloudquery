# Table: aws_lightsail_disk_snapshot


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_lightsail_disks`](aws_lightsail_disks.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|disk_arn|String|
|tags|JSON|
|arn|String|
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
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|