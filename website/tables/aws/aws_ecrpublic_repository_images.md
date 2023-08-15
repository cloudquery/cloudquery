# Table: aws_ecrpublic_repository_images

This table shows data for Ecrpublic Repository Images.

https://docs.aws.amazon.com/AmazonECRPublic/latest/APIReference/API_ImageDetail.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_ecrpublic_repositories](aws_ecrpublic_repositories).

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
|image_size_in_bytes|`int64`|
|image_tags|`list<item: utf8, nullable>`|
|registry_id|`utf8`|
|repository_name|`utf8`|