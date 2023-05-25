# Table: aws_iam_policies

This table shows data for IAM Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ManagedPolicyDetail.html

The composite primary key for this table is (**account_id**, **id**).

## Relations

The following tables depend on aws_iam_policies:
  - [aws_iam_policy_last_accessed_details](aws_iam_policy_last_accessed_details)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id (PK)|utf8|
|id (PK)|utf8|
|tags|json|
|policy_version_list|json|
|arn|utf8|
|attachment_count|int64|
|create_date|timestamp[us, tz=UTC]|
|default_version_id|utf8|
|description|utf8|
|is_attachable|bool|
|path|utf8|
|permissions_boundary_usage_count|int64|
|policy_id|utf8|
|policy_name|utf8|
|update_date|timestamp[us, tz=UTC]|