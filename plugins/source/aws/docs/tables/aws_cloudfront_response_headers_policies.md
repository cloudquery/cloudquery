# Table: aws_cloudfront_response_headers_policies

This table shows data for Cloudfront Response Headers Policies.

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListResponseHeadersPolicies.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|id|`utf8`|
|response_headers_policy|`json`|
|type|`utf8`|