# Table: aws_ec2_ebs_volumes

This table shows data for Amazon Elastic Compute Cloud (EC2) Amazon Elastic Block Store (EBS) Volumes.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Volume.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|attachments|`json`|
|availability_zone|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|encrypted|`bool`|
|fast_restored|`bool`|
|iops|`int64`|
|kms_key_id|`utf8`|
|multi_attach_enabled|`bool`|
|outpost_arn|`utf8`|
|size|`int64`|
|snapshot_id|`utf8`|
|sse_type|`utf8`|
|state|`utf8`|
|throughput|`int64`|
|volume_id|`utf8`|
|volume_type|`utf8`|