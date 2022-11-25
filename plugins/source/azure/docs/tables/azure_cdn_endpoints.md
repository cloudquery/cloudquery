# Table: azure_cdn_endpoints

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn#Endpoint

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
|cdn_profile_id|String|
|host_name|String|
|origins|JSON|
|origin_groups|JSON|
|resource_state|String|
|provisioning_state|String|
|origin_path|String|
|content_types_to_compress|StringArray|
|origin_host_header|String|
|is_compression_enabled|Bool|
|is_http_allowed|Bool|
|is_https_allowed|Bool|
|query_string_caching_behavior|String|
|optimization_type|String|
|probe_path|String|
|geo_filters|JSON|
|default_origin_group|JSON|
|url_signing_keys|JSON|
|delivery_policy|JSON|
|web_application_firewall_policy_link|JSON|
|location|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|
|system_data|JSON|