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
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
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