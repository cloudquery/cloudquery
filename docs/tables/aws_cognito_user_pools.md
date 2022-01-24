
# Table: aws_cognito_user_pools
A container for information about the user pool.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|account_recovery_setting|jsonb|Use this setting to define which verified available method a user can use to recover their password when they call ForgotPassword|
|admin_create_user_admin_only|boolean|Set to True if only the administrator is allowed to create user profiles|
|admin_create_user_invite_email_message|text|The message template for email messages|
|admin_create_user_invite_email_subject|text|The subject line for email messages|
|admin_create_user_invite_sms|text|The message template for SMS messages.|
|admin_create_user_config_unused_account_validity_days|integer|The user account expiration limit, in days, after which the account is no longer usable|
|alias_attributes|text[]|Specifies the attributes that are aliased in a user pool.|
|arn|text|The Amazon Resource Name (ARN) for the user pool.|
|auto_verified_attributes|text[]|Specifies the attributes that are auto-verified in a user pool.|
|creation_date|timestamp without time zone|The date the user pool was created.|
|custom_domain|text|A custom domain name that you provide to Amazon Cognito|
|challenge_required_on_new_device|boolean|Indicates whether a challenge is required on a new device|
|device_only_remembered_on_user_prompt|boolean|If true, a device is only remembered on user prompt.|
|domain|text|Holds the domain prefix if the user pool has a domain associated with it.|
|email_configuration_set|text|The set of configuration rules that can be applied to emails sent using Amazon SES|
|email_configuration_sending_account|text|Specifies whether Amazon Cognito emails your users by using its built-in email functionality or your Amazon SES email configuration|
|email_configuration_from|text|Identifies either the sender’s email address or the sender’s name with their email address|
|email_configuration_reply_to_address|text|The destination to which the receiver of the email should reply to.|
|email_configuration_source_arn|text|The Amazon Resource Name (ARN) of a verified email address in Amazon SES|
|email_configuration_failure|text|The reason why the email configuration cannot send the messages to your users.|
|email_verification_message|text|The contents of the email verification message.|
|email_verification_subject|text|The subject of the email verification message.|
|estimated_number_of_users|integer|A number estimating the size of the user pool.|
|id|text|The ID of the user pool.|
|lambda_config_create_auth_challenge|text|Creates an authentication challenge.|
|lambda_config_custom_email_sender_lambda_arn|text|The Lambda Amazon Resource Name of the Lambda function that Amazon Cognito triggers to send email notifications to users. |
|lambda_config_custom_email_sender_lambda_version|text|The Lambda version represents the signature of the "request" attribute in the "event" information Amazon Cognito passes to your custom email Lambda function. The only supported value is V1_0. |
|lambda_config_custom_message|text|A custom Message AWS Lambda trigger.|
|lambda_config_custom_sms_sender_lambda_arn|text|The Lambda Amazon Resource Name of the Lambda function that Amazon Cognito triggers to send SMS notifications to users. |
|lambda_config_custom_sms_sender_lambda_version|text|The Lambda version represents the signature of the "request" attribute in the "event" information Amazon Cognito passes to your custom SMS Lambda function. The only supported value is V1_0. |
|lambda_config_define_auth_challenge|text|Defines the authentication challenge.|
|lambda_config_kms_key_id|text|The Amazon Resource Name of Key Management Service Customer master keys|
|lambda_config_post_authentication|text|A post-authentication AWS Lambda trigger.|
|lambda_config_post_confirmation|text|A post-confirmation AWS Lambda trigger.|
|lambda_config_pre_authentication|text|A pre-authentication AWS Lambda trigger.|
|lambda_config_pre_sign_up|text|A pre-registration AWS Lambda trigger.|
|lambda_config_pre_token_generation|text|A Lambda trigger that is invoked before token generation.|
|lambda_config_user_migration|text|The user migration Lambda config type.|
|lambda_config_verify_auth_challenge_response|text|Verifies the authentication challenge response.|
|last_modified_date|timestamp without time zone|The date the user pool was last modified.|
|mfa_configuration|text|Can be one of the following values:  * OFF - MFA tokens are not required and cannot be specified during user registration.  * ON - MFA tokens are required for all user registrations|
|name|text|The name of the user pool.|
|policies_password_policy_minimum_length|integer|The minimum length of the password policy that you have set|
|policies_password_policy_require_lowercase|boolean|In the password policy that you have set, refers to whether you have required users to use at least one lowercase letter in their password.|
|policies_password_policy_require_numbers|boolean|In the password policy that you have set, refers to whether you have required users to use at least one number in their password.|
|policies_password_policy_require_symbols|boolean|In the password policy that you have set, refers to whether you have required users to use at least one symbol in their password.|
|policies_password_policy_require_uppercase|boolean|In the password policy that you have set, refers to whether you have required users to use at least one uppercase letter in their password.|
|policies_password_policy_temporary_password_validity_days|integer|In the password policy you have set, refers to the number of days a temporary password is valid|
|sms_authentication_message|text|The contents of the SMS authentication message.|
|sms_configuration_sns_caller_arn|text|The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) caller|
|sms_configuration_external_id|text|The external ID is a value that we recommend you use to add security to your IAM role which is used to call Amazon SNS to send SMS messages for your user pool. If you provide an ExternalId, the Cognito User Pool will include it when attempting to assume your IAM role, so that you can set your roles trust policy to require the ExternalID|
|sms_configuration_failure|text|The reason why the SMS configuration cannot send the messages to your users. This message might include comma-separated values to describe why your SMS configuration can't send messages to user pool end users.  * InvalidSmsRoleAccessPolicyException - The IAM role which Cognito uses to send SMS messages is not properly configured|
|sms_verification_message|text|The contents of the SMS verification message.|
|status|text|The status of a user pool.|
|user_pool_add_ons_advanced_security_mode|text|The advanced security mode. |
|user_pool_tags|jsonb|The tags that are assigned to the user pool|
|username_attributes|text[]|Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up.|
|username_configuration_case_sensitive|boolean|Specifies whether username case sensitivity will be applied for all users in the user pool through Cognito APIs|
|verification_message_template_default_email_option|text|The default email option.|
|verification_message_template_email_message|text|The email message template|
|verification_message_template_email_message_by_link|text|The email message template for sending a confirmation link to the user. EmailMessageByLink is allowed only if  EmailSendingAccount (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.|
|verification_message_template_email_subject|text|The subject line for the email message template|
|verification_message_template_email_subject_by_link|text|The subject line for the email message template for sending a confirmation link to the user|
|verification_message_template_sms_message|text|The SMS message template.|
