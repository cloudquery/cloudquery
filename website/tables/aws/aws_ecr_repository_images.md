# Table: aws_ecr_repository_images

This table shows data for Amazon Elastic Container Registry (ECR) Repository Images.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageDetail.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_ecr_repositories](aws_ecr_repositories).

The following tables depend on aws_ecr_repository_images:
  - [aws_ecr_repository_image_scan_findings](aws_ecr_repository_image_scan_findings)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|artifact_media_type|`utf8`|
|image_digest|`utf8`|
|image_manifest_media_type|`utf8`|
|image_pushed_at|`timestamp[us, tz=UTC]`|
|image_scan_findings_summary|`json`|
|image_scan_status|`json`|
|image_size_in_bytes|`int64`|
|image_tags|`list<item: utf8, nullable>`|
|last_recorded_pull_time|`timestamp[us, tz=UTC]`|
|registry_id|`utf8`|
|repository_name|`utf8`|

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


