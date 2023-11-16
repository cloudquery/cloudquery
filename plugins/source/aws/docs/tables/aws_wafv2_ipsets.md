# Table: aws_wafv2_ipsets

This table shows data for Wafv2 Ipsets.

https://docs.aws.amazon.com/waf/latest/APIReference/API_IPSet.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|addresses|`list<item: inet, nullable>`|
|tags|`json`|
|arn (PK)|`utf8`|
|ip_address_version|`utf8`|
|id|`utf8`|
|name|`utf8`|
|description|`utf8`|