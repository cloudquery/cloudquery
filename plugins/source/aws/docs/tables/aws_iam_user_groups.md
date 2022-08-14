
# Table: aws_iam_user_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_cq_id|uuid|Unique CloudQuery ID of aws_iam_users table (FK)|
|user_id|text|The stable and unique string identifying the user|
|group_arn|text|The Amazon Resource Name (ARN) specifying the group|
|create_date|timestamp without time zone|The date and time, in ISO 8601 date-time format, when the group was created|
|group_id|text|The stable and unique string identifying the group|
|group_name|text|The friendly name that identifies the group|
|path|text|The path to the group|
