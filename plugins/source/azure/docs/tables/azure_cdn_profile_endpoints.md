
# Table: azure_cdn_profile_endpoints
Endpoint CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content caching and delivery behavior
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|profile_cq_id|uuid|Unique CloudQuery ID of azure_cdn_profiles table (FK)|
|host_name|text|The host name of the endpoint structured as {endpointName}{DNSZone}, eg|
|resource_state|text|Resource status of the endpoint|
|provisioning_state|text|Provisioning status of the endpoint|
|origin_path|text|A directory path on the origin that CDN can use to retrieve content from, eg|
|content_types_to_compress|text[]|List of content types on which compression applies|
|origin_host_header|text|The host header value sent to the origin with each request|
|is_compression_enabled|boolean|Indicates whether content compression is enabled on CDN|
|is_http_allowed|boolean|Indicates whether HTTP traffic is allowed on the endpoint|
|is_https_allowed|boolean|Indicates whether HTTPS traffic is allowed on the endpoint|
|query_string_caching_behavior|text|Defines how CDN caches requests that include query strings|
|optimization_type|text|Specifies what scenario the customer wants this CDN endpoint to optimize for, eg|
|probe_path|text|Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN|
|default_origin_group_id|text|Resource ID|
|delivery_policy_description|text|User-friendly description of the policy|
|web_application_firewall_policy_link_id|text|Resource ID|
|location|text|Resource location|
|tags|jsonb|Resource tags|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
|created_by|text|An identifier for the identity that created the resource|
|created_by_type|text|The type of identity that created the resource|
|created_at_time|timestamp without time zone||
|last_modified_by|text|An identifier for the identity that last modified the resource|
|last_modified_by_type|text|The type of identity that last modified the resource|
|last_modified_at_time|timestamp without time zone||
