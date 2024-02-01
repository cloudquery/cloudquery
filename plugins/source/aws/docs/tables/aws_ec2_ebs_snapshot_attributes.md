# Table: aws_ec2_ebs_snapshot_attributes

This table shows data for Amazon Elastic Compute Cloud (EC2) Amazon Elastic Block Store (EBS) Snapshot Attributes.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshotAttribute.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **snapshot_arn**.
## Relations

This table depends on [aws_ec2_ebs_snapshots](aws_ec2_ebs_snapshots.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|snapshot_arn|`utf8`|
|create_volume_permissions|`json`|
|product_codes|`json`|
|snapshot_id|`utf8`|