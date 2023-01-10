# Table: aws_config_config_rules

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_config_config_rules:
  - [aws_config_config_rule_compliances](aws_config_config_rule_compliances.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|source|JSON|
|config_rule_arn|String|
|config_rule_id|String|
|config_rule_name|String|
|config_rule_state|String|
|created_by|String|
|description|String|
|evaluation_modes|JSON|
|input_parameters|String|
|maximum_execution_frequency|String|
|scope|JSON|