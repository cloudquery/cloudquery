# Table: aws_ssm_instance_compliance_items

This table shows data for AWS Systems Manager (SSM) Instance Compliance Items.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceItem.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**id**, **instance_arn**).
## Relations

This table depends on [aws_ssm_instances](aws_ssm_instances.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id|`utf8`|
|instance_arn|`utf8`|
|compliance_type|`utf8`|
|details|`json`|
|execution_summary|`json`|
|resource_id|`utf8`|
|resource_type|`utf8`|
|severity|`utf8`|
|status|`utf8`|
|title|`utf8`|