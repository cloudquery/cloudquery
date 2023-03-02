# Table: aws_ec2_launch_template_versions

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchTemplateVersion.html

The composite primary key for this table is (**arn**, **version_number**).

## Relations

This table depends on [aws_ec2_launch_templates](aws_ec2_launch_templates.md).

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
|version_number (PK)|Int|
|create_time|Timestamp|
|created_by|String|
|default_version|Bool|
|launch_template_data|JSON|
|launch_template_id|String|
|launch_template_name|String|
|version_description|String|