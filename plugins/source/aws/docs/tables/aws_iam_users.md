
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
|arn|text||
|password_enabled|boolean||
|password_last_changed|timestamp without time zone||
|password_next_rotation|timestamp without time zone||
|mfa_active|boolean||
|access_key1_active|boolean||
|access_key2_active|boolean||
|access_key1_last_rotated|timestamp without time zone||
|access_key2_last_rotated|timestamp without time zone||
|cert1_active|boolean||
|cert2_active|boolean||
|cert1_last_rotated|timestamp without time zone||
|cert2_last_rotated|timestamp without time zone||
