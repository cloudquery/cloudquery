# Table: aws_ecrpublic_repositories

https://docs.aws.amazon.com/AmazonECRPublic/latest/APIReference/API_Repository.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ecrpublic_repositories:
  - [aws_ecrpublic_repository_images](aws_ecrpublic_repository_images.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|created_at|Timestamp|
|registry_id|String|
|repository_name|String|
|repository_uri|String|