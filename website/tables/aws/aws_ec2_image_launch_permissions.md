# Table: aws_ec2_image_launch_permissions

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchPermission.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_ec2_images](aws_ec2_images).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|image_arn|String|
|group|String|
|organization_arn|String|
|organizational_unit_arn|String|
|user_id|String|