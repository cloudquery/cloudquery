# Table: aws_networkfirewall_firewall_policies

This table shows data for Networkfirewall Firewall Policies.

https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_FirewallPolicy.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|stateless_default_actions|`list<item: utf8, nullable>`|
|stateless_fragment_default_actions|`list<item: utf8, nullable>`|
|policy_variables|`json`|
|stateful_default_actions|`list<item: utf8, nullable>`|
|stateful_engine_options|`json`|
|stateful_rule_group_references|`json`|
|stateless_custom_actions|`json`|
|stateless_rule_group_references|`json`|
|tls_inspection_configuration_arn|`utf8`|
|firewall_policy_arn|`utf8`|
|firewall_policy_id|`utf8`|
|firewall_policy_name|`utf8`|
|consumed_stateful_rule_capacity|`int64`|
|consumed_stateless_rule_capacity|`int64`|
|description|`utf8`|
|encryption_configuration|`json`|
|firewall_policy_status|`utf8`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|number_of_associations|`int64`|