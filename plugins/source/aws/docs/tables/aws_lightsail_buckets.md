
# Table: aws_lightsail_buckets
Describes an Amazon Lightsail bucket
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|able_to_update_bundle|boolean|Indicates whether the bundle that is currently applied to a bucket can be changed to another bundle|
|access_log_config_enabled|boolean|A Boolean value that indicates whether bucket access logging is enabled for the bucket|
|access_log_config_destination|text|The name of the bucket where the access logs are saved|
|access_log_config_prefix|text|The optional object prefix for the bucket access log|
|access_rules_allow_public_overrides|boolean|A Boolean value that indicates whether the access control list (ACL) permissions that are applied to individual objects override the getObject option that is currently specified|
|access_rules_get_object|text|Specifies the anonymous access to all objects in a bucket|
|arn|text|The Amazon Resource Name (ARN) of the bucket|
|bundle_id|text|The ID of the bundle currently applied to the bucket|
|created_at|timestamp without time zone|The timestamp when the distribution was created|
|location_availability_zone|text|The Availability Zone|
|location_region_name|text|The AWS Region name|
|name|text|The name of the bucket|
|object_versioning|text|Indicates whether object versioning is enabled for the bucket|
|readonly_access_accounts|text[]|An array of strings that specify the Amazon Web Services account IDs that have read-only access to the bucket|
|resource_type|text|The Lightsail resource type of the bucket (for example, Bucket)|
|resources_receiving_access|jsonb|An array of objects that describe Lightsail instances that have access to the bucket|
|state_code|text|The state code of the bucket|
|state_message|text|A message that describes the state of the bucket|
|support_code|text|The support code for a bucket|
|tags|jsonb|The tag keys and optional values for the bucket|
|url|text|The URL of the bucket|
