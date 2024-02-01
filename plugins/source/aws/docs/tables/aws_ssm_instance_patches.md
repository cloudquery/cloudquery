# Table: aws_ssm_instance_patches

This table shows data for AWS Systems Manager (SSM) Instance Patches.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchComplianceData.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**instance_arn**, **kb_id**).
## Relations

This table depends on [aws_ssm_instances](aws_ssm_instances.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|instance_arn|`utf8`|
|kb_id|`utf8`|
|classification|`utf8`|
|installed_time|`timestamp[us, tz=UTC]`|
|severity|`utf8`|
|state|`utf8`|
|title|`utf8`|
|cve_ids|`utf8`|