# Table: aws_shield_attacks

This table shows data for Shield Attacks.

https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_AttackDetail.html

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|id (PK)|`utf8`|
|attack_counters|`json`|
|attack_id|`utf8`|
|attack_properties|`json`|
|end_time|`timestamp[us, tz=UTC]`|
|mitigations|`json`|
|resource_arn|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|sub_resources|`json`|