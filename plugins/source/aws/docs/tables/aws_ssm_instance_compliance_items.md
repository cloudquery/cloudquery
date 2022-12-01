# Table: aws_ssm_instance_compliance_items

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceItem.html

The composite primary key for this table is (**id**, **instance_arn**).

## Relations
This table depends on [aws_ssm_instances](aws_ssm_instances.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|id (PK)|String|
|instance_arn (PK)|String|
|compliance_type|String|
|details|JSON|
|execution_summary|JSON|
|resource_id|String|
|resource_type|String|
|severity|String|
|status|String|
|title|String|