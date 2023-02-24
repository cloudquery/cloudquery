# Table: aws_ssm_instance_patches

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchComplianceData.html

The composite primary key for this table is (**instance_arn**, **kb_id**).

## Relations

This table depends on [aws_ssm_instances](aws_ssm_instances).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|instance_arn (PK)|String|
|kb_id (PK)|String|
|classification|String|
|installed_time|Timestamp|
|severity|String|
|state|String|
|title|String|
|cve_ids|String|