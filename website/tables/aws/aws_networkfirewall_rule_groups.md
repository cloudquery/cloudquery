# Table: aws_networkfirewall_rule_groups

This table shows data for Networkfirewall Rule Groups.

https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_RuleGroup.html

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
|rules_source|JSON|
|reference_sets|JSON|
|rule_variables|JSON|
|stateful_rule_options|JSON|
|rule_group_arn|String|
|rule_group_id|String|
|rule_group_name|String|
|capacity|Int|
|consumed_capacity|Int|
|description|String|
|encryption_configuration|JSON|
|last_modified_time|Timestamp|
|number_of_associations|Int|
|rule_group_status|String|
|sns_topic|String|
|source_metadata|JSON|
|type|String|