# Table: aws_s3_bucket_grants

This table shows data for S3 Bucket Grants.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_Grant.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**bucket_arn**, **grantee_type**, **grantee_id**, **permission**).
## Relations

This table depends on [aws_s3_buckets](aws_s3_buckets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|bucket_arn|`utf8`|
|grantee_type|`utf8`|
|grantee_id|`utf8`|
|permission|`utf8`|
|grantee|`json`|