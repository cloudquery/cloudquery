# Table: aws_ssm_instances

This table shows data for AWS Systems Manager (SSM) Instances.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InstanceInformation.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ssm_instances:
  - [aws_ssm_instance_compliance_items](aws_ssm_instance_compliance_items)
  - [aws_ssm_instance_patches](aws_ssm_instance_patches)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|activation_id|`utf8`|
|agent_version|`utf8`|
|association_overview|`json`|
|association_status|`utf8`|
|computer_name|`utf8`|
|ip_address|`utf8`|
|iam_role|`utf8`|
|instance_id|`utf8`|
|is_latest_version|`bool`|
|last_association_execution_date|`timestamp[us, tz=UTC]`|
|last_ping_date_time|`timestamp[us, tz=UTC]`|
|last_successful_association_execution_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|ping_status|`utf8`|
|platform_name|`utf8`|
|platform_type|`utf8`|
|platform_version|`utf8`|
|registration_date|`timestamp[us, tz=UTC]`|
|resource_type|`utf8`|
|source_id|`utf8`|
|source_type|`utf8`|