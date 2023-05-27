# Table: gcp_storage_buckets

This table shows data for GCP Storage Buckets.

https://cloud.google.com/storage/docs/json_api/v1/buckets#resource

The primary key for this table is **name**.

## Relations

The following tables depend on gcp_storage_buckets:
  - [gcp_storage_bucket_policies](gcp_storage_bucket_policies)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|name (PK)|`utf8`|
|acl|`json`|
|bucket_policy_only|`json`|
|uniform_bucket_level_access|`json`|
|public_access_prevention|`int64`|
|default_object_acl|`json`|
|default_event_based_hold|`bool`|
|predefined_acl|`utf8`|
|predefined_default_object_acl|`utf8`|
|location|`utf8`|
|custom_placement_config|`json`|
|meta_generation|`int64`|
|storage_class|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|versioning_enabled|`bool`|
|labels|`json`|
|requester_pays|`bool`|
|lifecycle|`json`|
|retention_policy|`json`|
|cors|`json`|
|encryption|`json`|
|logging|`json`|
|website|`json`|
|etag|`utf8`|
|location_type|`utf8`|
|project_number|`int64`|
|rpo|`int64`|
|autoclass|`json`|