# Table: aws_s3_bucket_cors_rules

https://docs.aws.amazon.com/AmazonS3/latest/API/API_CORSRule.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_s3_buckets](aws_s3_buckets.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|bucket_arn|String|
|allowed_methods|StringArray|
|allowed_origins|StringArray|
|allowed_headers|StringArray|
|expose_headers|StringArray|
|id|String|
|max_age_seconds|Int|