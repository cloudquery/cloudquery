# Table: aws_s3_bucket_replications

This table shows data for S3 Bucket Replications.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketReplication.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **bucket_arn**.
## Relations

This table depends on [aws_s3_buckets](aws_s3_buckets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|bucket_arn|`utf8`|
|replication_configuration|`json`|