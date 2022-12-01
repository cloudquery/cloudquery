# Table: aws_ecrpublic_repository_images

https://docs.aws.amazon.com/AmazonECRPublic/latest/APIReference/API_ImageDetail.html

The primary key for this table is **arn**.

## Relations
This table depends on [aws_ecrpublic_repositories](aws_ecrpublic_repositories.md).


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
|image_size_in_bytes|Int|
|image_tags|StringArray|
|registry_id|String|
|repository_name|String|