# Table: aws_iam_policies


The composite primary key for this table is (**account_id**, **id**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|