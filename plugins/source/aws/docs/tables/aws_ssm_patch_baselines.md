# Table: aws_ssm_patch_baselines

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchBaselineIdentity.html

The composite primary key for this table is (**account_id**, **region**, **baseline_id**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|baseline_id (PK)|String|
|baseline_description|String|
|baseline_name|String|
|default_baseline|Bool|
|operating_system|String|