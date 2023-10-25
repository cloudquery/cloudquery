# Table: aws_ec2_image_launch_permissions

This table shows data for Amazon Elastic Compute Cloud (EC2) Image Launch Permissions.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchPermission.html

The primary key for this table is **image_arn**.

## Relations

This table depends on [aws_ec2_images](aws_ec2_images).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|image_arn (PK)|`utf8`|
|group|`utf8`|
|organization_arn|`utf8`|
|organizational_unit_arn|`utf8`|
|user_id|`utf8`|