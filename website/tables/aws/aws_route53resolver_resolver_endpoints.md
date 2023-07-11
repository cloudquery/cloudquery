# Table: aws_route53resolver_resolver_endpoints

This table shows data for Amazon Route 53 Resolver Resolver Endpoints.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ResolverEndpoint.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|creation_time|`utf8`|
|creator_request_id|`utf8`|
|direction|`utf8`|
|host_vpc_id|`utf8`|
|id|`utf8`|
|ip_address_count|`int64`|
|modification_time|`utf8`|
|name|`utf8`|
|resolver_endpoint_type|`utf8`|
|security_group_ids|`list<item: utf8, nullable>`|
|status|`utf8`|
|status_message|`utf8`|