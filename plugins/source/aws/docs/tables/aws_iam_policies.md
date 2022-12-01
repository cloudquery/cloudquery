# Table: aws_iam_policies

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ManagedPolicyDetail.html

The composite primary key for this table is (**account_id**, **id**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|id (PK)|String|
|tags|JSON|
|policy_version_list|JSON|
|arn|String|
|attachment_count|Int|
|create_date|Timestamp|
|default_version_id|String|
|description|String|
|is_attachable|Bool|
|path|String|
|permissions_boundary_usage_count|Int|
|policy_name|String|
|update_date|Timestamp|