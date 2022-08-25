
# Table: aws_iam_users

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|create_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the user was created.|
|path|text|The path to the user|
|id|text|The stable and unique string identifying the user|
|user_name|text|The friendly name identifying the user.|
|password_last_used|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the user's password was last used to sign in to an Amazon Web Services website|
|permissions_boundary_arn|text|The ARN of the policy used to set the permissions boundary for the user or role.|
|permissions_boundary_type|text|The permissions boundary usage type that indicates what type of IAM resource is used as the permissions boundary for an entity|
|tags|jsonb|A list of tags that are associated with the user|
|user|text||
|arn|text|The Amazon Resource Name (ARN) that identifies the user|
|password_enabled|boolean|When the user has a password, this value is TRUE|
|password_last_changed|timestamp without time zone|The date and time when the user's password was last set, in ISO 8601 date-time format|
|password_next_rotation|timestamp without time zone|When the account has a password policy that requires password rotation, this field contains the date and time, in ISO 8601 date-time format, when the user is required to set a new password|
|mfa_active|boolean|When a multi-factor authentication (MFA) device has been enabled for the user, this value is TRUE|
|access_key_1_active|boolean|When the user has an access key and the access key's status is Active, this value is TRUE|
|access_key_2_active|boolean|When the user has an access key and the access key's status is Active, this value is TRUE|
|access_key1_last_rotated|timestamp without time zone|The date and time, in ISO 8601 date-time format, when the user's access key was created or last changed|
|access_key2_last_rotated|timestamp without time zone|The date and time, in ISO 8601 date-time format, when the user's access key was created or last changed|
|cert_1_active|boolean|When the user has an X.509 signing certificate and that certificate's status is Active, this value is TRUE|
|cert_2_active|boolean|When the user has an X.509 signing certificate and that certificate's status is Active, this value is TRUE|
|cert_1_last_rotated|timestamp without time zone|The date and time, in ISO 8601 date-time format, when the user's signing certificate was created or last changed|
|cert_2_last_rotated|timestamp without time zone|The date and time, in ISO 8601 date-time format, when the user's signing certificate was created or last changed|
