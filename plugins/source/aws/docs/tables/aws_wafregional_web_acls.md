# Table: aws_wafregional_web_acls

This table shows data for Wafregional Web ACLs.

https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_WebACL.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|resources_for_web_acl|`list<item: utf8, nullable>`|
|default_action|`json`|
|rules|`json`|
|web_acl_id|`utf8`|
|metric_name|`utf8`|
|name|`utf8`|
|web_acl_arn|`utf8`|