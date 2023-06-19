# Table: aws_cloudfront_cache_policies

This table shows data for Cloudfront Cache Policies.

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CachePolicySummary.html

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|id (PK)|`utf8`|
|cache_policy|`json`|
|type|`utf8`|