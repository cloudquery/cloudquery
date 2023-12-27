# Table: aws_cloudfront_cache_policies

This table shows data for Cloudfront Cache Policies.

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CachePolicySummary.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|id|`utf8`|
|arn (PK)|`utf8`|
|cache_policy|`json`|
|type|`utf8`|