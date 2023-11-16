# Table: aws_networkfirewall_firewalls

This table shows data for Networkfirewall Firewalls.

https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_DescribeFirewall.html

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
|configuration_sync_state_summary|`utf8`|
|status|`utf8`|
|capacity_usage_summary|`json`|
|sync_states|`json`|
|firewall_id|`utf8`|
|firewall_policy_arn|`utf8`|
|subnet_mappings|`json`|
|vpc_id|`utf8`|
|delete_protection|`bool`|
|description|`utf8`|
|encryption_configuration|`json`|
|firewall_arn|`utf8`|
|firewall_name|`utf8`|
|firewall_policy_change_protection|`bool`|
|subnet_change_protection|`bool`|