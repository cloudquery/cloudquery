# Table: aws_cognito_identity_pools

The composite primary key for this table is (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|arn|String|
|id (PK)|String|
|saml_provider_ar_ns|StringArray|
|allow_unauthenticated_identities|Bool|
|identity_pool_id|String|
|identity_pool_name|String|
|allow_classic_flow|Bool|
|cognito_identity_providers|JSON|
|developer_provider_name|String|
|identity_pool_tags|JSON|
|open_id_connect_provider_ar_ns|StringArray|
|supported_login_providers|JSON|
|result_metadata|JSON|