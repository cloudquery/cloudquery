# Table: aws_waf_subscribed_rule_groups

This table shows data for WAF Subscribed Rule Groups.

https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_SubscribedRuleGroupSummary.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **rule_group_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|rule_group_id|`utf8`|
|metric_name|`utf8`|
|name|`utf8`|