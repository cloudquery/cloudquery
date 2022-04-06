
# Table: aws_iam_user_attached_policies

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_cq_id|uuid|Unique CloudQuery ID of aws_iam_users table (FK)|
|user_id|text|The stable and unique string identifying the user|
|policy_arn|text|The Amazon Resource Name (ARN) of the policy|
|policy_name|text|The friendly name of the attached policy|
