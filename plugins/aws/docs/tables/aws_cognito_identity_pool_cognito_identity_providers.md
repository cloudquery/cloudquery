
# Table: aws_cognito_identity_pool_cognito_identity_providers
A provider representing an Amazon Cognito user pool and its client ID.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|identity_pool_cq_id|uuid|Unique CloudQuery ID of aws_cognito_identity_pools table (FK)|
|identity_pool_id|text|An identity pool ID in the format REGION:GUID.|
|client_id|text|The client ID for the Amazon Cognito user pool.|
|provider_name|text|The provider name for an Amazon Cognito user pool|
|server_side_token_check|boolean|TRUE if server-side token validation is enabled for the identity provider’s token|
