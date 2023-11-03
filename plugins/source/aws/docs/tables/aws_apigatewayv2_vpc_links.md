# Table: aws_apigatewayv2_vpc_links

This table shows data for Amazon API Gateway v2 VPC Links.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/vpclinks.html#vpclinks-prop-vpclink

The composite primary key for this table is (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|name|`utf8`|
|security_group_ids|`list<item: utf8, nullable>`|
|subnet_ids|`list<item: utf8, nullable>`|
|vpc_link_id|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|tags|`json`|
|vpc_link_status|`utf8`|
|vpc_link_status_message|`utf8`|
|vpc_link_version|`utf8`|