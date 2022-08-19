
# Table: aws_iam_roles
An IAM role is an IAM identity that you can create in your account that has specific permissions.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|policies|jsonb|List of policies attached to group.|
|arn|text|The Amazon Resource Name (ARN) specifying the role|
|create_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the role was created.|
|path|text|The path to the role|
|id|text|The stable and unique string identifying the role|
|role_name|text|The friendly name that identifies the role.|
|assume_role_policy_document|jsonb|The policy that grants an entity permission to assume the role.|
|description|text|A description of the role that you provide.|
|max_session_duration|bigint|The maximum session duration (in seconds) for the specified role|
|permissions_boundary_arn|text|The ARN of the policy used to set the permissions boundary for the user or role.|
|permissions_boundary_type|text|The permissions boundary usage type that indicates what type of IAM resource is used as the permissions boundary for an entity|
|role_last_used_last_used_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601) that the role was last used|
|role_last_used_region|text|The name of the Amazon Web Services Region in which the role was last used.|
|tags|jsonb|A list of tags that are attached to the role|
