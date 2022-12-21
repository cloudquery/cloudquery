# Table: aws_amp_workspaces

https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference.html#AMP-APIReference-WorkspaceDescription

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_amp_workspaces:
  - [aws_amp_rule_groups_namespaces](aws_amp_rule_groups_namespaces.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|alert_manager_definition|JSON|
|logging_configuration|JSON|
|arn (PK)|String|
|created_at|Timestamp|
|status|JSON|
|workspace_id|String|
|alias|String|
|prometheus_endpoint|String|
|tags|JSON|