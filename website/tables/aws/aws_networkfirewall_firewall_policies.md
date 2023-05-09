# Table: aws_networkfirewall_firewall_policies

This table shows data for Networkfirewall Firewall Policies.

https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_FirewallPolicy.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|stateless_default_actions|StringArray|
|stateless_fragment_default_actions|StringArray|
|stateful_default_actions|StringArray|
|stateful_engine_options|JSON|
|stateful_rule_group_references|JSON|
|stateless_custom_actions|JSON|
|stateless_rule_group_references|JSON|
|tls_inspection_configuration_arn|String|
|firewall_policy_arn|String|
|firewall_policy_id|String|
|firewall_policy_name|String|
|consumed_stateful_rule_capacity|Int|
|consumed_stateless_rule_capacity|Int|
|description|String|
|encryption_configuration|JSON|
|firewall_policy_status|String|
|last_modified_time|Timestamp|
|number_of_associations|Int|