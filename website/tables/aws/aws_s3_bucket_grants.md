# Table: aws_s3_bucket_grants

This table shows data for S3 Bucket Grants.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_Grant.html

The composite primary key for this table is (**bucket_arn**, **grantee_type**, **grantee_id**, **permission**).

## Relations

This table depends on [aws_s3_buckets](aws_s3_buckets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|bucket_arn (PK)|utf8|
|grantee_type (PK)|utf8|
|grantee_id (PK)|utf8|
|permission (PK)|utf8|
|grantee|json|