# Table: aws_ec2_ebs_volume_statuses

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VolumeStatusItem.html

The primary key for this table is **volume_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|volume_arn (PK)|String|
|actions|JSON|
|attachment_statuses|JSON|
|availability_zone|String|
|events|JSON|
|outpost_arn|String|
|volume_id|String|
|volume_status|JSON|