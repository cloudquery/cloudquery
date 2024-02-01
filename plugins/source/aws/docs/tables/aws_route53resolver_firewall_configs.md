# Table: aws_route53resolver_firewall_configs

This table shows data for Amazon Route 53 Resolver Firewall Configs.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_FirewallConfig.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|firewall_fail_open|`utf8`|
|id|`utf8`|
|owner_id|`utf8`|
|resource_id|`utf8`|