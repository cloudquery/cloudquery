# Table: aws_route53resolver_firewall_rule_group_associations

This table shows data for Amazon Route 53 Resolver Firewall Rule Group Associations.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_FirewallRuleGroupAssociation.html

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
|firewall_rule_group_id|`utf8`|
|id|`utf8`|
|managed_owner_name|`utf8`|
|modification_time|`utf8`|
|mutation_protection|`utf8`|
|name|`utf8`|
|priority|`int64`|
|status|`utf8`|
|status_message|`utf8`|
|vpc_id|`utf8`|