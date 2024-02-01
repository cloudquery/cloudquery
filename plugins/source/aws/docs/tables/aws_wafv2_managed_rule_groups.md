# Table: aws_wafv2_managed_rule_groups

This table shows data for Wafv2 Managed Rule Groups.

https://docs.aws.amazon.com/waf/latest/APIReference/API_ManagedRuleGroupSummary.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **scope**, **name**, **vendor_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|scope|`utf8`|
|properties|`json`|
|description|`utf8`|
|name|`utf8`|
|vendor_name|`utf8`|
|versioning_supported|`bool`|