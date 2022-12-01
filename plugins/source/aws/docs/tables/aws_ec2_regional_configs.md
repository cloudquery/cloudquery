# Table: aws_ec2_regional_configs



The composite primary key for this table is (**account_id**, **region**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|ebs_encryption_enabled_by_default|Bool|
|ebs_default_kms_key_id|String|