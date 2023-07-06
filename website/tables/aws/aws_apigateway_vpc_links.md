# Table: aws_apigateway_vpc_links

This table shows data for Amazon API Gateway VPC Links.

https://docs.aws.amazon.com/apigateway/latest/api/API_VpcLink.html

The composite primary key for this table is (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|description|`utf8`|
|id|`utf8`|
|name|`utf8`|
|status|`utf8`|
|status_message|`utf8`|
|tags|`json`|
|target_arns|`list<item: utf8, nullable>`|