
# Table: aws_cognito_identity_pools
An object representing an Amazon Cognito identity pool.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|allow_unauthenticated_identities|boolean|TRUE if the identity pool supports unauthenticated logins.|
|id|text|An identity pool ID in the format REGION:GUID.|
|identity_pool_name|text|A string that you provide.|
|allow_classic_flow|boolean|Enables or disables the Basic (Classic) authentication flow|
|developer_provider_name|text|The "domain" by which Cognito will refer to your users.|
|identity_pool_tags|jsonb|The tags that are assigned to the identity pool|
|open_id_connect_provider_arns|text[]|The ARNs of the OpenID Connect providers.|
|saml_provider_arns|text[]|An array of Amazon Resource Names (ARNs) of the SAML provider for your identity pool.|
|supported_login_providers|jsonb|Optional key:value pairs mapping provider names to provider app IDs.|
