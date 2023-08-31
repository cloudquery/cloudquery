# Table: aws_s3_bucket_websites

This table shows data for S3 Bucket Websites.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_WebsiteConfiguration.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_s3_buckets](aws_s3_buckets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|bucket_arn|`utf8`|
|error_document|`json`|
|index_document|`json`|
|redirect_all_requests_to|`json`|
|routing_rules|`json`|