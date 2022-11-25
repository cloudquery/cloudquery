# Table: aws_wafregional_rate_based_rules

https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_RateBasedRule.html

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
|match_predicates|JSON|
|rate_key|String|
|rate_limit|Int|
|rule_id|String|
|metric_name|String|
|name|String|