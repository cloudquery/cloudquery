# Table: aws_wafv2_rule_groups

This table shows data for Wafv2 Rule Groups.

https://docs.aws.amazon.com/waf/latest/APIReference/API_RuleGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|policy|`json`|
|capacity|`int64`|
|id|`utf8`|
|name|`utf8`|
|visibility_config|`json`|
|available_labels|`json`|
|consumed_labels|`json`|
|custom_response_bodies|`json`|
|description|`utf8`|
|label_namespace|`utf8`|
|rules|`json`|