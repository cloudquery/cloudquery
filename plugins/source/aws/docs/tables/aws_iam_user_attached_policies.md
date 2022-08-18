
# Table: aws_iam_user_attached_policies
Contains information about an attached policy
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_cq_id|uuid|Unique CloudQuery ID of aws_iam_users table (FK)|
|policy_arn|text|The Amazon Resource Name (ARN)|
|policy_name|text|The friendly name of the attached policy.|
