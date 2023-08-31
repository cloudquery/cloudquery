# Table: aws_shield_protection_groups

This table shows data for Shield Protection Groups.

https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_ProtectionGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|aggregation|`utf8`|
|members|`list<item: utf8, nullable>`|
|pattern|`utf8`|
|protection_group_id|`utf8`|
|protection_group_arn|`utf8`|
|resource_type|`utf8`|