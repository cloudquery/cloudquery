# Table: aws_cloudfront_response_headers_policies

This table shows data for Cloudfront Response Headers Policies.

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListResponseHeadersPolicies.html

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
|response_headers_policy|`json`|
|type|`utf8`|