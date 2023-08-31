# Table: aws_waf_rule_groups

This table shows data for WAF Rule Groups.

https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_RuleGroupSummary.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|rule_ids|`list<item: utf8, nullable>`|
|rule_group_id|`utf8`|
|metric_name|`utf8`|
|name|`utf8`|