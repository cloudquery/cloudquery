# Table: aws_lightsail_distributions

This table shows data for Lightsail Distributions.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetDistributions.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|able_to_update_bundle|`bool`|
|alternative_domain_names|`list<item: utf8, nullable>`|
|bundle_id|`utf8`|
|cache_behavior_settings|`json`|
|cache_behaviors|`json`|
|certificate_name|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|default_cache_behavior|`json`|
|domain_name|`utf8`|
|ip_address_type|`utf8`|
|is_enabled|`bool`|
|location|`json`|
|name|`utf8`|
|origin|`json`|
|origin_public_dns|`utf8`|
|resource_type|`utf8`|
|status|`utf8`|
|support_code|`utf8`|
|latest_cache_reset|`json`|