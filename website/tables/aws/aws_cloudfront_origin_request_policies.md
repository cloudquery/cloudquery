# Table: aws_cloudfront_origin_request_policies

This table shows data for Cloudfront Origin Request Policies.

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListOriginRequestPolicies.html

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|id (PK)|`utf8`|
|origin_request_policy|`json`|
|type|`utf8`|