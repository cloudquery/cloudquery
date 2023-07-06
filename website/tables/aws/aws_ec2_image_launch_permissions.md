# Table: aws_ec2_image_launch_permissions

This table shows data for Amazon Elastic Compute Cloud (EC2) Image Launch Permissions.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchPermission.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_ec2_images](aws_ec2_images).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|image_arn|`utf8`|
|group|`utf8`|
|organization_arn|`utf8`|
|organizational_unit_arn|`utf8`|
|user_id|`utf8`|