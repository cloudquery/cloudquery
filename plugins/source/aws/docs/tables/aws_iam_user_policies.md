
# Table: aws_iam_user_policies
Contains the response to a successful GetUserPolicy request.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_cq_id|uuid|Unique CloudQuery ID of aws_iam_users table (FK)|
|policy_document|text|The policy document|
|policy_name|text|The name of the policy.  This member is required.|
|user_name|text|The user the policy is associated with.  This member is required.|
