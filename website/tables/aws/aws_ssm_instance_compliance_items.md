# Table: aws_ssm_instance_compliance_items

This table shows data for AWS Systems Manager (SSM) Instance Compliance Items.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceItem.html

The composite primary key for this table is (**id**, **instance_arn**).

## Relations

This table depends on [aws_ssm_instances](aws_ssm_instances).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|id (PK)|utf8|
|instance_arn (PK)|utf8|
|compliance_type|utf8|
|details|json|
|execution_summary|json|
|resource_id|utf8|
|resource_type|utf8|
|severity|utf8|
|status|utf8|
|title|utf8|