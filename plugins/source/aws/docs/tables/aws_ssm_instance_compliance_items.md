# Table: aws_ssm_instance_compliance_items


The composite primary key for this table is (**id**, **instance_arn**).

## Relations
This table depends on [`aws_ssm_instances`](aws_ssm_instances.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|