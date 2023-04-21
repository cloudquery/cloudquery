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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|data_encryption_key_id|String|
|description|String|
|encrypted|Bool|
|kms_key_id|String|
|outpost_arn|String|
|owner_alias|String|
|owner_id|String|
|progress|String|
|restore_expiry_time|Timestamp|
|snapshot_id|String|
|start_time|Timestamp|
|state|String|
|state_message|String|
|storage_tier|String|
|volume_id|String|
|volume_size|Int|