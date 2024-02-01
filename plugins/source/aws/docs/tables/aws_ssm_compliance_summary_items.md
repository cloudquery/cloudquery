# Table: aws_ssm_compliance_summary_items

This table shows data for AWS Systems Manager (SSM) Compliance Summary Items.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceSummaryItem.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **compliance_type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|compliance_type|`utf8`|
|compliant_summary|`json`|
|non_compliant_summary|`json`|