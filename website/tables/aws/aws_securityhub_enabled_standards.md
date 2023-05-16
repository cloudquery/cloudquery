# Table: aws_securityhub_enabled_standards

This table shows data for AWS Security Hub Enabled Standards.

https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_GetEnabledStandards.html

The composite primary key for this table is (**account_id**, **region**, **standards_arn**, **standards_subscription_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|standards_arn (PK)|String|
|standards_input|JSON|
|standards_status|String|
|standards_subscription_arn (PK)|String|
|standards_status_reason|JSON|