# Table: aws_wafregional_rate_based_rules

This table shows data for Wafregional Rate Based Rules.

https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_RateBasedRule.html

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
|match_predicates|`json`|
|rate_key|`utf8`|
|rate_limit|`int64`|
|rule_id|`utf8`|
|metric_name|`utf8`|
|name|`utf8`|