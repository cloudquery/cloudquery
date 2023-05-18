# Table: aws_efs_access_points

This table shows data for Amazon Elastic File System (EFS) Access Points.

https://docs.aws.amazon.com/efs/latest/ug/API_AccessPointDescription.html

The primary key for this table is **arn**.

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
|tags|JSON|
|access_point_arn|String|
|access_point_id|String|
|client_token|String|
|file_system_id|String|
|life_cycle_state|String|
|name|String|
|owner_id|String|
|posix_user|JSON|
|root_directory|JSON|