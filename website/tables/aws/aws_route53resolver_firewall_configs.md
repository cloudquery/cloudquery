# Table: aws_route53resolver_firewall_configs

This table shows data for Amazon Route 53 Resolver Firewall Configs.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_FirewallConfig.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|firewall_fail_open|`utf8`|
|id (PK)|`utf8`|
|owner_id|`utf8`|
|resource_id|`utf8`|