# Table: aws_amp_workspaces

This table shows data for Amazon Managed Service for Prometheus (AMP) Workspaces.

https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference.html#AMP-APIReference-WorkspaceDescription

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_amp_workspaces:
  - [aws_amp_rule_groups_namespaces](aws_amp_rule_groups_namespaces)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|alert_manager_definition|`json`|
|logging_configuration|`json`|
|arn (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|status|`json`|
|workspace_id|`utf8`|
|alias|`utf8`|
|prometheus_endpoint|`utf8`|
|tags|`json`|