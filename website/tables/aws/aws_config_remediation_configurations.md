# Table: aws_config_remediation_configurations

This table shows data for Config Remediation Configurations.

https://docs.aws.amazon.com/config/latest/APIReference/API_RemediationConfiguration.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_config_config_rules](aws_config_config_rules).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|config_rule_name|`utf8`|
|target_id|`utf8`|
|target_type|`utf8`|
|arn (PK)|`utf8`|
|automatic|`bool`|
|created_by_service|`utf8`|
|execution_controls|`json`|
|maximum_automatic_attempts|`int64`|
|parameters|`json`|
|resource_type|`utf8`|
|retry_attempt_seconds|`int64`|
|target_version|`utf8`|