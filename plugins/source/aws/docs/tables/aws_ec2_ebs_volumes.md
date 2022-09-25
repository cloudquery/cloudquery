# Table: aws_ec2_ebs_volumes


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|attachments|JSON|
|availability_zone|String|
|create_time|Timestamp|
|encrypted|Bool|
|fast_restored|Bool|
|iops|Int|
|kms_key_id|String|
|multi_attach_enabled|Bool|
|outpost_arn|String|
|size|Int|
|snapshot_id|String|
|state|String|
|tags|JSON|
|throughput|Int|
|volume_id|String|
|volume_type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|