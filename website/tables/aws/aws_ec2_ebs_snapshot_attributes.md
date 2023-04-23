# Table: aws_ec2_ebs_snapshot_attributes

This table shows data for Amazon Elastic Compute Cloud (EC2) Amazon Elastic Block Store (EBS) Snapshot Attributes.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshotAttribute.html

The primary key for this table is **snapshot_arn**.

## Relations

This table depends on [aws_ec2_ebs_snapshots](aws_ec2_ebs_snapshots).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|snapshot_arn (PK)|String|
|create_volume_permissions|JSON|
|product_codes|JSON|
|snapshot_id|String|