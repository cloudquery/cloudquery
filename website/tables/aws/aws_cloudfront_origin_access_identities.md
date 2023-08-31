# Table: aws_cloudfront_origin_access_identities

This table shows data for Cloudfront Origin Access Identities.

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListCloudFrontOriginAccessIdentities.html

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|comment|`utf8`|
|id (PK)|`utf8`|
|s3_canonical_user_id|`utf8`|