# Table: aws_amp_rule_groups_namespaces

This table shows data for Amazon Managed Service for Prometheus (AMP) Rule Groups Namespaces.

https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference.html#AMP-APIReference-RuleGroupsNamespaceDescription

The primary key for this table is **arn**.

## Relations

This table depends on [aws_amp_workspaces](aws_amp_workspaces).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|workspace_arn|`utf8`|
|arn (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|data|`binary`|
|modified_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|status|`json`|
|tags|`json`|