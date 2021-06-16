
# Table: aws_iam_role_policies
Inline policies that are embedded in the specified IAM role
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|role_id|uuid|Role ID the policy belongs too.|
|account_id|text|The AWS Account ID of the resource.|
|policy_document|jsonb|The policy document. IAM stores policies in JSON format. However, resources that were created using AWS CloudFormation templates can be formatted in YAML. AWS CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.|
|policy_name|text|The name of the policy.|
|role_name|text|The role the policy is associated with.|
