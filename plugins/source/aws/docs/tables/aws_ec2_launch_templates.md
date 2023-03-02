# Table: aws_ec2_launch_templates

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchTemplate.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ec2_launch_templates:
  - [aws_ec2_launch_template_versions](aws_ec2_launch_template_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|create_time|Timestamp|
|created_by|String|
|default_version_number|Int|
|latest_version_number|Int|
|launch_template_id|String|
|launch_template_name|String|