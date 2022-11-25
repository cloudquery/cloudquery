# Table: aws_ec2_ebs_volumes

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Volume.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
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