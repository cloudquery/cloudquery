# Table: azure_cdn_routes

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn#Route

The primary key for this table is **id**.

## Relations
This table depends on [azure_cdn_endpoints](azure_cdn_endpoints.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|cache_configuration|JSON|
|custom_domains|JSON|
|enabled_state|String|
|forwarding_protocol|String|
|https_redirect|String|
|link_to_default_domain|String|
|origin_group|JSON|
|origin_path|String|
|patterns_to_match|StringArray|
|rule_sets|JSON|
|supported_protocols|StringArray|
|deployment_status|String|
|endpoint_name|String|
|provisioning_state|String|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|
|endpoint_id|String|