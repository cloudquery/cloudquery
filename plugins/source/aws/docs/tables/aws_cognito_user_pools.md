# Table: aws_cognito_user_pools

https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolType.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Relations

The following tables depend on aws_cognito_user_pools:
  - [aws_cognito_user_pool_identity_providers](aws_cognito_user_pool_identity_providers.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|id (PK)|String|
|account_recovery_setting|JSON|
|admin_create_user_config|JSON|
|alias_attributes|StringArray|
|arn|String|
|auto_verified_attributes|StringArray|
|creation_date|Timestamp|
|custom_domain|String|
|deletion_protection|String|
|device_configuration|JSON|
|domain|String|
|email_configuration|JSON|
|email_configuration_failure|String|
|email_verification_message|String|
|email_verification_subject|String|
|estimated_number_of_users|Int|
|lambda_config|JSON|
|last_modified_date|Timestamp|
|mfa_configuration|String|
|name|String|
|policies|JSON|
|schema_attributes|JSON|
|sms_authentication_message|String|
|sms_configuration|JSON|
|sms_configuration_failure|String|
|sms_verification_message|String|
|status|String|
|user_attribute_update_settings|JSON|
|user_pool_add_ons|JSON|
|user_pool_tags|JSON|
|username_attributes|StringArray|
|username_configuration|JSON|
|verification_message_template|JSON|