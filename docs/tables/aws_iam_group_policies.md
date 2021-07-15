
# Table: aws_iam_group_policies
Inline policies that are embedded in the specified IAM group
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|group_cq_id|uuid|Unique CloudQuery ID of aws_iam_groups table (FK)|
|group_id|text|Group ID the policy belongs too.|
|group_name|text|The group the policy is associated with.|
|policy_document|jsonb|The policy document. IAM stores policies in JSON format. However, resources that were created using AWS CloudFormation templates can be formatted in YAML. AWS CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.|
|policy_name|text|The name of the policy.|
