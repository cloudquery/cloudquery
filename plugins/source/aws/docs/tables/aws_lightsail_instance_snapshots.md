# Table: aws_lightsail_instance_snapshots

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_InstanceSnapshot.html

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
|created_at|Timestamp|
|from_attached_disks|JSON|
|from_blueprint_id|String|
|from_bundle_id|String|
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
|tags|JSON|