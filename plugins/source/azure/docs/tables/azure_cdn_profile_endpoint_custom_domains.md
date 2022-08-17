
# Table: azure_cdn_profile_endpoint_custom_domains
CustomDomain friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, eg
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|profile_endpoint_cq_id|uuid|Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)|
|host_name|text|The host name of the custom domain|
|state|text|Resource status of the custom domain|
|custom_https_provisioning_state|text|Provisioning status of Custom Https of the custom domain|
|custom_https_provisioning_substate|text|Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step|
|custom_https_parameters|jsonb|Certificate parameters for securing custom HTTPS|
|validation_data|text|Special validation or data may be required when delivering CDN to some regions due to local compliance reasons|
|provisioning_state|text|Provisioning status of the custom domain|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
|created_by|text|An identifier for the identity that created the resource|
|created_by_type|text|The type of identity that created the resource|
|created_at_time|timestamp without time zone||
|last_modified_by|text|An identifier for the identity that last modified the resource|
|last_modified_by_type|text|The type of identity that last modified the resource|
|last_modified_at_time|timestamp without time zone||
