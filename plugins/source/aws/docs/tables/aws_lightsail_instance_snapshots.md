# Table: aws_lightsail_instance_snapshots

This table shows data for Lightsail Instance Snapshots.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_InstanceSnapshot.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|from_attached_disks|`json`|
|from_blueprint_id|`utf8`|
|from_bundle_id|`utf8`|
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