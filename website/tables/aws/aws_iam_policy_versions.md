# Table: aws_iam_policy_versions

This table shows data for IAM Policy Versions.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_PolicyVersion.html

The composite primary key for this table is (**account_id**, **policy_arn**).

## Relations

This table depends on [aws_iam_policies](aws_iam_policies).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|policy_arn (PK)|`utf8`|
|policy_version|`json`|
|result_metadata|`json`|