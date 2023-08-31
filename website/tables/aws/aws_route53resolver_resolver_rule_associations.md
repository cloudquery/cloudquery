# Table: aws_route53resolver_resolver_rule_associations

This table shows data for Amazon Route 53 Resolver Resolver Rule Associations.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ResolverRuleAssociation.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|resolver_rule_id|`utf8`|
|status|`utf8`|
|status_message|`utf8`|
|vpc_id|`utf8`|