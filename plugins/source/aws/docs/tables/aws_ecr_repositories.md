# Table: aws_ecr_repositories



The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_ecr_repositories`:
  - [`aws_ecr_repository_images`](aws_ecr_repository_images.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|created_at|Timestamp|
|encryption_configuration|JSON|
|image_scanning_configuration|JSON|
|image_tag_mutability|String|
|registry_id|String|
|repository_name|String|
|repository_uri|String|