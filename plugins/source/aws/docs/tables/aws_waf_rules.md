# Table: aws_waf_rules

This table shows data for WAF Rules.

https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_RuleSummary.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|tags|`json`|
|predicates|`json`|
|rule_id|`utf8`|
|metric_name|`utf8`|
|name|`utf8`|