# Table: aws_ssm_compliance_summary_items

This table shows data for AWS Systems Manager (SSM) Compliance Summary Items.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceSummaryItem.html

The composite primary key for this table is (**account_id**, **region**, **compliance_type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|compliance_type (PK)|`utf8`|
|compliant_summary|`json`|
|non_compliant_summary|`json`|