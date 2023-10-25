# Table: aws_ec2_image_last_launched_times

This table shows data for Amazon Elastic Compute Cloud (EC2) Image Last Launched Times.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeImageAttribute.html. 
The date and time, in ISO 8601 date-time format, when the AMI was last used to launch an EC2 instance. When the AMI is used to launch an instance, there is a 24-hour delay before that usage is reported.

The primary key for this table is **image_arn**.

## Relations

This table depends on [aws_ec2_images](aws_ec2_images).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|image_arn (PK)|`utf8`|
|last_launched_time|`timestamp[us, tz=UTC]`|