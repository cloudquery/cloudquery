# Table: aws_networkfirewall_rule_groups

This table shows data for Networkfirewall Rule Groups.

https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_RuleGroup.html

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
|rules_source|`json`|
|reference_sets|`json`|
|rule_variables|`json`|
|stateful_rule_options|`json`|
|rule_group_arn|`utf8`|
|rule_group_id|`utf8`|
|rule_group_name|`utf8`|
|capacity|`int64`|
|consumed_capacity|`int64`|
|description|`utf8`|
|encryption_configuration|`json`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|number_of_associations|`int64`|
|rule_group_status|`utf8`|
|sns_topic|`utf8`|
|source_metadata|`json`|
|type|`utf8`|