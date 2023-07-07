# Table: aws_ec2_ebs_snapshots

This table shows data for Amazon Elastic Compute Cloud (EC2) Amazon Elastic Block Store (EBS) Snapshots.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Snapshot.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ec2_ebs_snapshots:
  - [aws_ec2_ebs_snapshot_attributes](aws_ec2_ebs_snapshot_attributes)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|data_encryption_key_id|`utf8`|
|description|`utf8`|
|encrypted|`bool`|
|kms_key_id|`utf8`|
|outpost_arn|`utf8`|
|owner_alias|`utf8`|
|owner_id|`utf8`|
|progress|`utf8`|
|restore_expiry_time|`timestamp[us, tz=UTC]`|
|snapshot_id|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|state_message|`utf8`|
|storage_tier|`utf8`|
|volume_id|`utf8`|
|volume_size|`int64`|