# Table: aws_cognito_user_pool_identity_providers

This table shows data for Cognito User Pool Identity Providers.

https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_IdentityProviderType.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_cognito_user_pools](aws_cognito_user_pools).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|user_pool_arn|`utf8`|
|attribute_mapping|`json`|
|creation_date|`timestamp[us, tz=UTC]`|
|idp_identifiers|`list<item: utf8, nullable>`|
|last_modified_date|`timestamp[us, tz=UTC]`|
|provider_details|`json`|
|provider_name|`utf8`|
|provider_type|`utf8`|
|user_pool_id|`utf8`|