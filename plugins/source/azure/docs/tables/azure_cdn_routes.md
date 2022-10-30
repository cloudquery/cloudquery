# Table: azure_cdn_routes



The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|cdn_endpoint_id|String|
|custom_domains|JSON|
|origin_group|JSON|
|origin_path|String|
|rule_sets|JSON|
|supported_protocols|StringArray|
|patterns_to_match|StringArray|
|query_string_caching_behavior|String|
|forwarding_protocol|String|
|link_to_default_domain|String|
|https_redirect|String|
|enabled_state|String|
|provisioning_state|String|
|deployment_status|String|
|id (PK)|String|
|name|String|
|type|String|
|system_data|JSON|