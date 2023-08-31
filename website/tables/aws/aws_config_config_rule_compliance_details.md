# Table: aws_config_config_rule_compliance_details

This table shows data for Config Config Rule Compliance Details.

https://docs.aws.amazon.com/config/latest/APIReference/API_EvaluationResult.html

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
|config_rule_name|`utf8`|
|annotation|`utf8`|
|compliance_type|`utf8`|
|config_rule_invoked_time|`timestamp[us, tz=UTC]`|
|evaluation_result_identifier|`json`|
|result_recorded_time|`timestamp[us, tz=UTC]`|
|result_token|`utf8`|