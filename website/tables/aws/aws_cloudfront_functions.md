# Table: aws_cloudfront_functions

This table shows data for Cloudfront Functions.

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_DescribeFunction.html

The composite primary key for this table is (**stage**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|stage (PK)|`utf8`|
|arn (PK)|`utf8`|
|e_tag|`utf8`|
|function_summary|`json`|
|result_metadata|`json`|