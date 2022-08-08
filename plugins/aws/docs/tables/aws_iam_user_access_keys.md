
# Table: aws_iam_user_access_keys

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_cq_id|uuid|Unique CloudQuery ID of aws_iam_users table (FK)|
|user_id|text|The stable and unique string identifying the user|
|access_key_id|text|The ID for this access key|
|create_date|timestamp without time zone|The date when the access key was created|
|status|text|The status of the access key. Active means that the key is valid for API calls; Inactive means it is not|
|last_used|timestamp without time zone|The date and time, in ISO 8601 date-time format, when the user's second access key was most recently used to sign an AWS API request|
|last_rotated|timestamp without time zone|The date and time, in ISO 8601 date-time format, when the user's access key was created or last changed|
|last_used_service_name|text|The AWS service that was most recently accessed with the user's second access key|
