# Table: aws_config_config_rule_compliances

https://docs.aws.amazon.com/config/latest/APIReference/API_ComplianceByConfigRule.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_config_config_rules](aws_config_config_rules.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|compliance|JSON|
|config_rule_name|String|