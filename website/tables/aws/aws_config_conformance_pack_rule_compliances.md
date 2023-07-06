# Table: aws_config_conformance_pack_rule_compliances

This table shows data for Config Conformance Pack Rule Compliances.

https://docs.aws.amazon.com/config/latest/APIReference/API_DescribeConformancePackCompliance.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_config_conformance_packs](aws_config_conformance_packs).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|conformance_pack_arn|`utf8`|
|compliance_type|`utf8`|
|config_rule_name|`utf8`|
|controls|`list<item: utf8, nullable>`|
|config_rule_invoked_time|`timestamp[us, tz=UTC]`|
|evaluation_result_identifier|`json`|
|result_recorded_time|`timestamp[us, tz=UTC]`|
|annotation|`utf8`|