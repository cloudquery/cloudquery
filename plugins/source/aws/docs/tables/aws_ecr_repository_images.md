# Table: aws_ecr_repository_images

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageDetail.html

The primary key for this table is **arn**.

## Relations
This table depends on [aws_ecr_repositories](aws_ecr_repositories.md).

The following tables depend on aws_ecr_repository_images:
  - [aws_ecr_repository_image_scan_findings](aws_ecr_repository_image_scan_findings.md)

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
|artifact_media_type|String|
|image_digest|String|
|image_manifest_media_type|String|
|image_pushed_at|Timestamp|
|image_scan_findings_summary|JSON|
|image_scan_status|JSON|
|image_size_in_bytes|Int|
|image_tags|StringArray|
|last_recorded_pull_time|Timestamp|
|registry_id|String|
|repository_name|String|