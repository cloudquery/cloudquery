# Table: aws_ec2_ebs_snapshot_attributes

This table shows data for Amazon Elastic Compute Cloud (EC2) Amazon Elastic Block Store (EBS) Snapshot Attributes.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshotAttribute.html

The primary key for this table is **snapshot_arn**.

## Relations

This table depends on [aws_ec2_ebs_snapshots](aws_ec2_ebs_snapshots).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|snapshot_arn (PK)|`utf8`|
|create_volume_permissions|`json`|
|product_codes|`json`|
|snapshot_id|`utf8`|