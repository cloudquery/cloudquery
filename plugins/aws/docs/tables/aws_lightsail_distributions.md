
# Table: aws_lightsail_distributions

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|able_to_update_bundle|boolean|Indicates whether the bundle that is currently applied to your distribution, specified using the distributionName parameter, can be changed to another bundle|
|alternative_domain_names|text[]|The alternate domain names of the distribution|
|arn|text|The Amazon Resource Name (ARN) of the distribution|
|bundle_id|text|The ID of the bundle currently applied to the distribution|
|cache_behavior_settings|jsonb|An object that describes the cache behavior settings of the distribution|
|cache_behaviors|jsonb|An array of objects that describe the per-path cache behavior of the distribution|
|certificate_name|text|The name of the SSL/TLS certificate attached to the distribution, if any|
|created_at|timestamp without time zone|The timestamp when the distribution was created|
|default_cache_behavior|text|The cache behavior of the distribution|
|domain_name|text|The domain name of the distribution|
|ip_address_type|text|The IP address type of the distribution|
|is_enabled|boolean|Indicates whether the distribution is enabled|
|availability_zone|text|The Availability Zone|
|name|text|The name of the distribution|
|origin_name|text|The name of the origin resource|
|origin_protocol_policy|text|The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content|
|origin_region_name|text|The AWS Region name of the origin resource|
|origin_resource_type|text|The resource type of the origin resource (eg, Instance)|
|origin_public_dns|text|The public DNS of the origin|
|resource_type|text|The Lightsail resource type (eg, Distribution)|
|status|text|The status of the distribution|
|support_code|text|The support code|
|tags|jsonb|The tag keys and optional values for the resource|
|cache_reset_create_time|timestamp without time zone|The timestamp of the last cache reset (eg, 147973490917) in Unix time format|
|cache_reset_status|text|The status of the last cache reset|
