# Table: aws_cloudfront_cache_policies

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CachePolicySummary.html

The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|id|String|
|arn|String|
|cache_policy|JSON|
|type|String|