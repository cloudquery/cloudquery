
# Table: aws_iam_role_policies
Contains the response to a successful GetRolePolicy request.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|role_cq_id|uuid|Unique CloudQuery ID of aws_iam_roles table (FK)|
|policy_document|jsonb|The policy document|
|policy_name|text|The name of the policy.|
|role_name|text|The role the policy is associated with.|
