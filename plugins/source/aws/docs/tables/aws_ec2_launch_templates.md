# Table: aws_ec2_launch_templates

This table shows data for Amazon Elastic Compute Cloud (EC2) Launch Templates.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchTemplate.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_ec2_launch_templates:
  - [aws_ec2_launch_template_versions](aws_ec2_launch_template_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|created_by|`utf8`|
|default_version_number|`int64`|
|latest_version_number|`int64`|
|launch_template_id|`utf8`|
|launch_template_name|`utf8`|