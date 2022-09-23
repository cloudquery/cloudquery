# Table: aws_lightsail_buckets


The primary key for this table is **_cq_id**.

## Relations
The following tables depend on `aws_lightsail_buckets`:
  - [`aws_lightsail_bucket_access_keys`](aws_lightsail_bucket_access_keys.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|tags|JSON|
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
|url|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|