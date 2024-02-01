# Table: aws_ecr_repository_images

This table shows data for Amazon Elastic Container Registry (ECR) Repository Images.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageDetail.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**repository_arn**, **image_digest**, **registry_id**).
## Relations

This table depends on [aws_ecr_repositories](aws_ecr_repositories.md).

The following tables depend on aws_ecr_repository_images:
  - [aws_ecr_repository_image_scan_findings](aws_ecr_repository_image_scan_findings.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|repository_arn|`utf8`|
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