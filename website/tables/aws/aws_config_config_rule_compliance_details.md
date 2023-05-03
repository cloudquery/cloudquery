# Table: aws_config_config_rule_compliance_details

This table shows data for Config Config Rule Compliance Details.

https://docs.aws.amazon.com/config/latest/APIReference/API_EvaluationResult.html

The composite primary key for this table is (**account_id**, **region**, **_evaluation_hash**).

## Relations

This table depends on [aws_config_config_rules](aws_config_config_rules).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|config_rule_name|String|
|_evaluation_hash (PK)|String|
|annotation|String|
|compliance_type|String|
|config_rule_invoked_time|Timestamp|
|evaluation_result_identifier|JSON|
|result_recorded_time|Timestamp|
|result_token|String|