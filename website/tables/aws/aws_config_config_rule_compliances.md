# Table: aws_config_config_rule_compliances

This table shows data for Config Config Rule Compliances.

https://docs.aws.amazon.com/config/latest/APIReference/API_ComplianceByConfigRule.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_config_config_rules](aws_config_config_rules).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|compliance|`json`|
|config_rule_name|`utf8`|