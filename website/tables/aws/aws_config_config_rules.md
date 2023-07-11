# Table: aws_config_config_rules

This table shows data for Config Config Rules.

https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigRule.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_config_config_rules:
  - [aws_config_config_rule_compliance_details](aws_config_config_rule_compliance_details)
  - [aws_config_config_rule_compliances](aws_config_config_rule_compliances)
  - [aws_config_remediation_configurations](aws_config_remediation_configurations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|source|`json`|
|config_rule_arn|`utf8`|
|config_rule_id|`utf8`|
|config_rule_name|`utf8`|
|config_rule_state|`utf8`|
|created_by|`utf8`|
|description|`utf8`|
|evaluation_modes|`json`|
|input_parameters|`utf8`|
|maximum_execution_frequency|`utf8`|
|scope|`json`|