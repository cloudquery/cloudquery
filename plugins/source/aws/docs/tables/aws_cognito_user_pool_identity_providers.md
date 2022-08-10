
# Table: aws_cognito_user_pool_identity_providers
A container for information about an identity provider.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_pool_cq_id|uuid|Unique CloudQuery ID of aws_cognito_user_pools table (FK)|
|user_pool_id|text|The ID of the user pool.|
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|attribute_mapping|jsonb|A mapping of identity provider attributes to standard and custom user pool attributes.|
|creation_date|timestamp without time zone|The date the identity provider was created.|
|idp_identifiers|text[]|A list of identity provider identifiers.|
|last_modified_date|timestamp without time zone|The date the identity provider was last modified.|
|provider_details|jsonb|The identity provider details|
|provider_name|text|The identity provider name.|
|provider_type|text|The identity provider type.|
