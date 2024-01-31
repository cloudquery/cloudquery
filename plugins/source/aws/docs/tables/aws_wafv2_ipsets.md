# Table: aws_wafv2_ipsets

This table shows data for Wafv2 Ipsets.

https://docs.aws.amazon.com/waf/latest/APIReference/API_IPSet.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|addresses|`list<item: inet, nullable>`|
|tags|`json`|
|arn|`utf8`|
|ip_address_version|`utf8`|
|id|`utf8`|
|name|`utf8`|
|description|`utf8`|