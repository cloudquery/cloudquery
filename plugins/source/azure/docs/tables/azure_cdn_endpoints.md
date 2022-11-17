# Table: azure_cdn_endpoints

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn#Endpoint

The primary key for this table is **id**.

## Relations
This table depends on [azure_cdn_profiles](azure_cdn_profiles.md).

The following tables depend on azure_cdn_endpoints:
  - [azure_cdn_custom_domains](azure_cdn_custom_domains.md)
  - [azure_cdn_routes](azure_cdn_routes.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|origins|JSON|
|content_types_to_compress|StringArray|
|default_origin_group|JSON|
|delivery_policy|JSON|
|geo_filters|JSON|
|is_compression_enabled|Bool|
|is_http_allowed|Bool|
|is_https_allowed|Bool|
|optimization_type|String|
|origin_groups|JSON|
|origin_host_header|String|
|origin_path|String|
|probe_path|String|
|query_string_caching_behavior|String|
|url_signing_keys|JSON|
|web_application_firewall_policy_link|JSON|
|custom_domains|JSON|
|host_name|String|
|provisioning_state|String|
|resource_state|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|
|profile_id|String|