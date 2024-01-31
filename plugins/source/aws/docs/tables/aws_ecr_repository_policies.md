# Table: aws_ecr_repository_policies

This table shows data for Amazon Elastic Container Registry (ECR) Repository Policies.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_GetRepositoryPolicy.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**repository_arn**, **registry_id**).
## Relations

This table depends on [aws_ecr_repositories](aws_ecr_repositories.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|repository_arn|`utf8`|
|policy_json|`json`|
|policy_text|`utf8`|
|registry_id|`utf8`|
|repository_name|`utf8`|