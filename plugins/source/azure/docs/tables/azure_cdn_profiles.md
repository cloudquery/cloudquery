
# Table: azure_cdn_profiles
Profile CDN profile is a logical grouping of endpoints that share the same settings, such as CDN provider and pricing tier
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription ID|
|sku_name|text|Name of the pricing tier|
|state|text|Resource status of the profile|
|provisioning_state|text|Provisioning status of the profile|
|frontdoor_id|text|The Id of the frontdoor|
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
