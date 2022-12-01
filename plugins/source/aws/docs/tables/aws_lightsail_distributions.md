# Table: aws_lightsail_distributions



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
|tags|JSON|
|able_to_update_bundle|Bool|
|alternative_domain_names|StringArray|
|bundle_id|String|
|cache_behavior_settings|JSON|
|cache_behaviors|JSON|
|certificate_name|String|
|created_at|Timestamp|
|default_cache_behavior|JSON|
|domain_name|String|
|ip_address_type|String|
|is_enabled|Bool|
|location|JSON|
|name|String|
|origin|JSON|
|origin_public_dns|String|
|resource_type|String|
|status|String|
|support_code|String|
|latest_cache_reset|JSON|