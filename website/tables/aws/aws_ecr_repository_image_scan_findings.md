# Table: aws_ecr_repository_image_scan_findings

This table shows data for Amazon Elastic Container Registry (ECR) Repository Image Scan Findings.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageScanFindings.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_ecr_repository_images](aws_ecr_repository_images).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|image_tag|String|
|image_digest|String|
|image_scan_findings|JSON|
|image_scan_status|JSON|
|registry_id|String|
|repository_name|String|