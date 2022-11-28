# Table: aws_ssm_instances

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InstanceInformation.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ssm_instances:
  - [aws_ssm_instance_compliance_items](aws_ssm_instance_compliance_items.md)
  - [aws_ssm_instance_patches](aws_ssm_instance_patches.md)

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
|activation_id|String|
|agent_version|String|
|association_overview|JSON|
|association_status|String|
|computer_name|String|
|ip_address|String|
|iam_role|String|
|instance_id|String|
|is_latest_version|Bool|
|last_association_execution_date|Timestamp|
|last_ping_date_time|Timestamp|
|last_successful_association_execution_date|Timestamp|
|name|String|
|ping_status|String|
|platform_name|String|
|platform_type|String|
|platform_version|String|
|registration_date|Timestamp|
|resource_type|String|
|source_id|String|
|source_type|String|