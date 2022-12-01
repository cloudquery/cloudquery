# Table: aws_ssm_compliance_summary_items

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceSummaryItem.html

The composite primary key for this table is (**account_id**, **region**, **compliance_type**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|compliance_type (PK)|String|
|compliant_summary|JSON|
|non_compliant_summary|JSON|