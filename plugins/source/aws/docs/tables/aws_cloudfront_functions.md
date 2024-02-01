# Table: aws_cloudfront_functions

This table shows data for Cloudfront Functions.

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_DescribeFunction.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**stage**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|stage|`utf8`|
|arn|`utf8`|
|e_tag|`utf8`|
|function_summary|`json`|