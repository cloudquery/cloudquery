# Table: aws_ssoadmin_permission_sets

This table shows data for Ssoadmin Permission Sets.

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_PermissionSet.html. 
The 'request_account_id' and 'request_region' columns are added to show the account_id and region of where the request was made from.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_ssoadmin_instances](aws_ssoadmin_instances).

The following tables depend on aws_ssoadmin_permission_sets:
  - [aws_ssoadmin_account_assignments](aws_ssoadmin_account_assignments)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|request_account_id|String|
|request_region|String|
|inline_policy|JSON|
|created_date|Timestamp|
|description|String|
|name|String|
|permission_set_arn|String|
|relay_state|String|
|session_duration|String|