# Table: aws_waf_rule_groups

This table shows data for WAF Rule Groups.

https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_RuleGroupSummary.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|tags|JSON|
|rule_ids|StringArray|
|rule_group_id|String|
|metric_name|String|
|name|String|