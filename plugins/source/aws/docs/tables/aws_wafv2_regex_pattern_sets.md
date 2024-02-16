# Table: aws_wafv2_regex_pattern_sets

This table shows data for Wafv2 Regex Pattern Sets.

https://docs.aws.amazon.com/waf/latest/APIReference/API_RegexPatternSet.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|description|`utf8`|
|id|`utf8`|
|name|`utf8`|
|regular_expression_list|`json`|