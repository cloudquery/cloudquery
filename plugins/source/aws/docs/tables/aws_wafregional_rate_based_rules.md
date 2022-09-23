# Table: aws_wafregional_rate_based_rules


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|