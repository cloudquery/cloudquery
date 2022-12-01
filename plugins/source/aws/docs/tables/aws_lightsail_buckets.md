# Table: aws_lightsail_buckets

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Bucket.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_lightsail_buckets:
  - [aws_lightsail_bucket_access_keys](aws_lightsail_bucket_access_keys.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|able_to_update_bundle|Bool|
|access_log_config|JSON|
|access_rules|JSON|
|arn|String|
|bundle_id|String|
|created_at|Timestamp|
|location|JSON|
|name|String|
|object_versioning|String|
|readonly_access_accounts|StringArray|
|resource_type|String|
|resources_receiving_access|JSON|
|state|JSON|
|support_code|String|
|tags|JSON|
|url|String|