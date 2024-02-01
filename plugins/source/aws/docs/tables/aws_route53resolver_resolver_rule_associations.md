# Table: aws_route53resolver_resolver_rule_associations

This table shows data for Amazon Route 53 Resolver Resolver Rule Associations.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ResolverRuleAssociation.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id|`utf8`|
|name|`utf8`|
|resolver_rule_id|`utf8`|
|status|`utf8`|
|status_message|`utf8`|
|vpc_id|`utf8`|