# Table: aws_cognito_user_pools

This table shows data for Cognito User Pools.

https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolType.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Relations

The following tables depend on aws_cognito_user_pools:
  - [aws_cognito_user_pool_identity_providers](aws_cognito_user_pool_identity_providers)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|id (PK)|`utf8`|
|account_recovery_setting|`json`|
|admin_create_user_config|`json`|
|alias_attributes|`list<item: utf8, nullable>`|
|arn|`utf8`|
|auto_verified_attributes|`list<item: utf8, nullable>`|
|creation_date|`timestamp[us, tz=UTC]`|
|custom_domain|`utf8`|
|deletion_protection|`utf8`|
|device_configuration|`json`|
|domain|`utf8`|
|email_configuration|`json`|
|email_configuration_failure|`utf8`|
|email_verification_message|`utf8`|
|email_verification_subject|`utf8`|
|estimated_number_of_users|`int64`|
|lambda_config|`json`|
|last_modified_date|`timestamp[us, tz=UTC]`|
|mfa_configuration|`utf8`|
|name|`utf8`|
|policies|`json`|
|schema_attributes|`json`|
|sms_authentication_message|`utf8`|
|sms_configuration|`json`|
|sms_configuration_failure|`utf8`|
|sms_verification_message|`utf8`|
|status|`utf8`|
|user_attribute_update_settings|`json`|
|user_pool_add_ons|`json`|
|user_pool_tags|`json`|
|username_attributes|`list<item: utf8, nullable>`|
|username_configuration|`json`|
|verification_message_template|`json`|