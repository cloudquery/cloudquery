# Table: aws_ecr_repository_lifecycle_policies

This table shows data for Amazon Elastic Container Registry (ECR) Repository Lifecycle Policies.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_GetLifecyclePolicy.html

The composite primary key for this table is (**account_id**, **region**, **registry_id**, **repository_name**).

## Relations

This table depends on [aws_ecr_repositories](aws_ecr_repositories.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|policy_json|`json`|
|last_evaluated_at|`timestamp[us, tz=UTC]`|
|lifecycle_policy_text|`utf8`|
|registry_id (PK)|`utf8`|
|repository_name (PK)|`utf8`|