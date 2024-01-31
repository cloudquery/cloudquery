# Table: aws_cloudfront_cache_policies

This table shows data for Cloudfront Cache Policies.

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CachePolicySummary.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|id|`utf8`|
|arn|`utf8`|
|cache_policy|`json`|
|type|`utf8`|