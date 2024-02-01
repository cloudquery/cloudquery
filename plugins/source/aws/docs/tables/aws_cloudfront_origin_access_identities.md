# Table: aws_cloudfront_origin_access_identities

This table shows data for Cloudfront Origin Access Identities.

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListCloudFrontOriginAccessIdentities.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|comment|`utf8`|
|id|`utf8`|
|s3_canonical_user_id|`utf8`|