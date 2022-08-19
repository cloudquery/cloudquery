
# Table: aws_iam_user_groups
Contains information about an IAM group entity
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_cq_id|uuid|Unique CloudQuery ID of aws_iam_users table (FK)|
|arn|text|The Amazon Resource Name (ARN) specifying the group|
|create_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the group was created.|
|group_id|text|The stable and unique string identifying the group|
|group_name|text|The friendly name that identifies the group.|
|path|text|The path to the group|
