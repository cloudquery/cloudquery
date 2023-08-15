# Table: aws_cognito_identity_pools

This table shows data for Cognito Identity Pools.

https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_DescribeIdentityPool.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|arn|`utf8`|
|id (PK)|`utf8`|
|saml_provider_arns|`list<item: utf8, nullable>`|
|allow_unauthenticated_identities|`bool`|
|identity_pool_id|`utf8`|
|identity_pool_name|`utf8`|
|allow_classic_flow|`bool`|
|cognito_identity_providers|`json`|
|developer_provider_name|`utf8`|
|identity_pool_tags|`json`|
|open_id_connect_provider_arns|`list<item: utf8, nullable>`|
|supported_login_providers|`json`|
|result_metadata|`json`|