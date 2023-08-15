# Table: aws_elbv2_load_balancer_web_acls

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Load Balancer Web ACLs.

https://docs.aws.amazon.com/waf/latest/APIReference/API_GetWebACLForResource.html

The composite primary key for this table is (**load_balancer_arn**, **arn**).

## Relations

This table depends on [aws_elbv2_load_balancers](aws_elbv2_load_balancers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|load_balancer_arn (PK)|`utf8`|
|arn (PK)|`utf8`|
|default_action|`json`|
|id|`utf8`|
|name|`utf8`|
|visibility_config|`json`|
|association_config|`json`|
|capacity|`int64`|
|captcha_config|`json`|
|challenge_config|`json`|
|custom_response_bodies|`json`|
|description|`utf8`|
|label_namespace|`utf8`|
|managed_by_firewall_manager|`bool`|
|post_process_firewall_manager_rule_groups|`json`|
|pre_process_firewall_manager_rule_groups|`json`|
|rules|`json`|
|token_domains|`list<item: utf8, nullable>`|