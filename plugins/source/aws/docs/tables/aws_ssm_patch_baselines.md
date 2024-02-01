# Table: aws_ssm_patch_baselines

This table shows data for AWS Systems Manager (SSM) Patch Baselines.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchBaselineIdentity.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **baseline_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|baseline_id|`utf8`|
|baseline_description|`utf8`|
|baseline_name|`utf8`|
|default_baseline|`bool`|
|operating_system|`utf8`|