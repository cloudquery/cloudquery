# Table: aws_wafv2_managed_rule_groups

This table shows data for Wafv2 Managed Rule Groups.

https://docs.aws.amazon.com/waf/latest/APIReference/API_ManagedRuleGroupSummary.html

The composite primary key for this table is (**account_id**, **region**, **scope**, **name**, **vendor_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|scope (PK)|`utf8`|
|properties|`json`|
|description|`utf8`|
|name (PK)|`utf8`|
|vendor_name (PK)|`utf8`|
|versioning_supported|`bool`|