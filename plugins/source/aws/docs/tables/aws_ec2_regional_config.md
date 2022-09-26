# Table: aws_ec2_regional_config
Ec2 Regional Config defines common default configuration for ec2 service

The composite primary key for this table is (**account_id**, **region**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id (PK)|String|
|region (PK)|String|
|ebs_encryption_enabled_by_default|Bool|
|ebs_default_kms_key_id|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|