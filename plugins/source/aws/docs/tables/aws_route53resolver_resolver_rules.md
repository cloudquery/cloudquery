# Table: aws_route53resolver_resolver_rules

This table shows data for Amazon Route 53 Resolver Resolver Rules.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ResolverRule.html

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
|domain_name|`utf8`|
|id|`utf8`|
|modification_time|`utf8`|
|name|`utf8`|
|owner_id|`utf8`|
|resolver_endpoint_id|`utf8`|
|rule_type|`utf8`|
|share_status|`utf8`|
|status|`utf8`|
|status_message|`utf8`|
|target_ips|`json`|