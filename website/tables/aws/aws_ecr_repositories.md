# Table: aws_ecr_repositories

This table shows data for Amazon Elastic Container Registry (ECR) Repositories.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_Repository.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ecr_repositories:
  - [aws_ecr_repository_images](aws_ecr_repository_images)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|tags|json|
|policy_text|json|
|created_at|timestamp[us, tz=UTC]|
|encryption_configuration|json|
|image_scanning_configuration|json|
|image_tag_mutability|utf8|
|registry_id|utf8|
|repository_arn|utf8|
|repository_name|utf8|
|repository_uri|utf8|