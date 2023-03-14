# Table: aws_s3_bucket_grants

This table shows data for AWS S3 Bucket Grants.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_Grant.html

The composite primary key for this table is (**bucket_arn**, **grantee_type**, **grantee_id**, **permission**).

## Relations

This table depends on [aws_s3_buckets](aws_s3_buckets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|bucket_arn (PK)|String|
|grantee_type (PK)|String|
|grantee_id (PK)|String|
|permission (PK)|String|
|grantee|JSON|