# Table: aws_ecr_repository_image_scan_findings

This table shows data for Amazon Elastic Container Registry (ECR) Repository Image Scan Findings.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageScanFindings.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_ecr_repository_images](aws_ecr_repository_images).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|image_tag|`utf8`|
|image_digest|`utf8`|
|image_scan_findings|`json`|
|image_scan_status|`json`|
|registry_id|`utf8`|
|repository_name|`utf8`|