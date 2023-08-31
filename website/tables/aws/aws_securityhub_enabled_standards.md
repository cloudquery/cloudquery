# Table: aws_securityhub_enabled_standards

This table shows data for AWS Security Hub Enabled Standards.

https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_GetEnabledStandards.html

The composite primary key for this table is (**account_id**, **region**, **standards_arn**, **standards_subscription_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|standards_arn (PK)|`utf8`|
|standards_input|`json`|
|standards_status|`utf8`|
|standards_subscription_arn (PK)|`utf8`|
|standards_status_reason|`json`|