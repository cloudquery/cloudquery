# Table: aws_lightsail_buckets

This table shows data for Lightsail Buckets.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Bucket.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_lightsail_buckets:
  - [aws_lightsail_bucket_access_keys](aws_lightsail_bucket_access_keys.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|able_to_update_bundle|`bool`|
|tags|`json`|
|access_log_config|`json`|
|access_rules|`json`|
|arn|`utf8`|
|bundle_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|location|`json`|
|name|`utf8`|
|object_versioning|`utf8`|
|readonly_access_accounts|`list<item: utf8, nullable>`|
|resource_type|`utf8`|
|resources_receiving_access|`json`|
|state|`json`|
|support_code|`utf8`|
|url|`utf8`|