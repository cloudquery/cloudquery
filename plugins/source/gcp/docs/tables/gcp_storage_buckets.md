# Table: gcp_storage_buckets



The primary key for this table is **name**.

## Relations

The following tables depend on gcp_storage_buckets:
  - [gcp_storage_bucket_policies](gcp_storage_bucket_policies.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|name (PK)|String|
|acl|JSON|
|bucket_policy_only|JSON|
|uniform_bucket_level_access|JSON|
|public_access_prevention|Int|
|default_object_acl|JSON|
|default_event_based_hold|Bool|
|predefined_acl|String|
|predefined_default_object_acl|String|
|location|String|
|custom_placement_config|JSON|
|meta_generation|Int|
|storage_class|String|
|created|Timestamp|
|versioning_enabled|Bool|
|labels|JSON|
|requester_pays|Bool|
|lifecycle|JSON|
|retention_policy|JSON|
|cors|JSON|
|encryption|JSON|
|logging|JSON|
|website|JSON|
|etag|String|
|location_type|String|
|project_number|Int|
|rpo|Int|
|autoclass|JSON|