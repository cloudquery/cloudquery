# Table: aws_iam_policy_versions

This table shows data for IAM Policy Versions.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_PolicyVersion.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **policy_arn**, **version_id**).
## Relations

This table depends on [aws_iam_policies](aws_iam_policies.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|policy_arn|`utf8`|
|document_json|`json`|
|create_date|`timestamp[us, tz=UTC]`|
|document|`utf8`|
|is_default_version|`bool`|
|version_id|`utf8`|