# Table: aws_ecr_repository_image_scan_findings

This table shows data for Amazon Elastic Container Registry (ECR) Repository Image Scan Findings.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageScanFindings.html

The composite primary key for this table is (**repository_arn**, **image_digest**).

## Relations

This table depends on [aws_ecr_repository_images](aws_ecr_repository_images.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|repository_arn (PK)|`utf8`|
|image_digest (PK)|`utf8`|
|image_id|`json`|
|image_scan_findings|`json`|
|image_scan_status|`json`|
|registry_id|`utf8`|
|repository_name|`utf8`|