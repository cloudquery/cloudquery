# Table: aws_ec2_ebs_volume_statuses

This table shows data for Amazon Elastic Compute Cloud (EC2) Amazon Elastic Block Store (EBS) Volume Statuses.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VolumeStatusItem.html

The primary key for this table is **volume_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|volume_arn (PK)|`utf8`|
|actions|`json`|
|attachment_statuses|`json`|
|availability_zone|`utf8`|
|events|`json`|
|outpost_arn|`utf8`|
|volume_id|`utf8`|
|volume_status|`json`|