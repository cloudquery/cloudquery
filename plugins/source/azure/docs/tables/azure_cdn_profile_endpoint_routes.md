
# Table: azure_cdn_profile_endpoint_routes
Route friendly Routes name mapping to the any Routes or secret related information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|profile_endpoint_cq_id|uuid|Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)|
|custom_domains|text[]|Domains referenced by this endpoint|
|origin_group_id|text|Resource ID|
|origin_path|text|A directory path on the origin that AzureFrontDoor can use to retrieve content from, eg|
|rule_sets|text[]|rule sets referenced by this endpoint|
|supported_protocols|text[]|List of supported protocols for this route|
|patterns_to_match|text[]|The route patterns of the rule|
|compression_settings|jsonb|compression settings|
|query_string_caching_behavior|text|Defines how CDN caches requests that include query strings|
|forwarding_protocol|text|Protocol this rule will use when forwarding traffic to backends|
|link_to_default_domain|text|whether this route will be linked to the default endpoint domain|
|https_redirect|text|Whether to automatically redirect HTTP traffic to HTTPS traffic|
|enabled_state|text|Whether to enable use of this rule|
|provisioning_state|text|Provisioning status|
|deployment_status|text|Possible values include: 'DeploymentStatusNotStarted', 'DeploymentStatusInProgress', 'DeploymentStatusSucceeded', 'DeploymentStatusFailed'|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
|created_by|text|An identifier for the identity that created the resource|
|created_by_type|text|The type of identity that created the resource|
|created_at_time|timestamp without time zone||
|last_modified_by|text|An identifier for the identity that last modified the resource|
|last_modified_by_type|text|The type of identity that last modified the resource|
|last_modified_at_time|timestamp without time zone||
