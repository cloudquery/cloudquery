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
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|policy_text|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|encryption_configuration|`json`|
|image_scanning_configuration|`json`|
|image_tag_mutability|`utf8`|
|registry_id|`utf8`|
|repository_arn|`utf8`|
|repository_name|`utf8`|
|repository_uri|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Unused ECR repository

```sql
WITH
  image
    AS (
      SELECT DISTINCT account_id, repository_name FROM aws_ecr_repository_images
    )
SELECT
  'Unused ECR repository' AS title,
  repository.account_id,
  repository.arn AS resource_id,
  'fail' AS status
FROM
  aws_ecr_repositories AS repository
  LEFT JOIN image ON
      image.account_id = repository.account_id
      AND image.repository_name = repository.repository_name
WHERE
  image.repository_name IS NULL;
```


