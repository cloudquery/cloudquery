# Table: aws_route53resolver_firewall_domain_lists

This table shows data for Amazon Route 53 Resolver Firewall Domain Lists.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_FirewallDomainList.html

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
|domain_count|`int64`|
|id|`utf8`|
|managed_owner_name|`utf8`|
|modification_time|`utf8`|
|name|`utf8`|
|status|`utf8`|
|status_message|`utf8`|