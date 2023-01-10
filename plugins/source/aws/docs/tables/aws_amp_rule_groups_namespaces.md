# Table: aws_amp_rule_groups_namespaces

https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference.html#AMP-APIReference-RuleGroupsNamespaceDescription

The primary key for this table is **arn**.

## Relations

This table depends on [aws_amp_workspaces](aws_amp_workspaces.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|workspace_arn|String|
|arn (PK)|String|
|created_at|Timestamp|
|data|ByteArray|
|modified_at|Timestamp|
|name|String|
|status|JSON|
|tags|JSON|