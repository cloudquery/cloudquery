# Table: aws_ssm_instance_patches

This table shows data for AWS Systems Manager (SSM) Instance Patches.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchComplianceData.html

The composite primary key for this table is (**instance_arn**, **kb_id**).

## Relations

This table depends on [aws_ssm_instances](aws_ssm_instances).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|instance_arn (PK)|`utf8`|
|kb_id (PK)|`utf8`|
|classification|`utf8`|
|installed_time|`timestamp[us, tz=UTC]`|
|severity|`utf8`|
|state|`utf8`|
|title|`utf8`|
|cve_ids|`utf8`|