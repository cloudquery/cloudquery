# Table: aws_ec2_ebs_volumes

This table shows data for Amazon Elastic Compute Cloud (EC2) Amazon Elastic Block Store (EBS) Volumes.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Volume.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|tags|json|
|attachments|json|
|availability_zone|utf8|
|create_time|timestamp[us, tz=UTC]|
|encrypted|bool|
|fast_restored|bool|
|iops|int64|
|kms_key_id|utf8|
|multi_attach_enabled|bool|
|outpost_arn|utf8|
|size|int64|
|snapshot_id|utf8|
|state|utf8|
|throughput|int64|
|volume_id|utf8|
|volume_type|utf8|