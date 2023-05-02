# Table: aws_config_remediation_configurations

This table shows data for Config Remediation Configurations.

https://docs.aws.amazon.com/config/latest/APIReference/API_DescribeRemediationConfigurations.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_config_config_rules](aws_config_config_rules).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|config_rule_name|String|
|target_id|String|
|target_type|String|
|arn (PK)|String|
|automatic|Bool|
|created_by_service|String|
|execution_controls|JSON|
|maximum_automatic_attempts|Int|
|parameters|JSON|
|resource_type|String|
|retry_attempt_seconds|Int|
|target_version|String|