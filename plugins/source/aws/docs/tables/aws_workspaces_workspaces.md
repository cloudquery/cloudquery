# Table: aws_workspaces_workspaces

https://docs.aws.amazon.com/workspaces/latest/api/API_Workspace.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|bundle_id|String|
|computer_name|String|
|directory_id|String|
|error_code|String|
|error_message|String|
|ip_address|String|
|modification_states|JSON|
|related_workspaces|JSON|
|root_volume_encryption_enabled|Bool|
|state|String|
|subnet_id|String|
|user_name|String|
|user_volume_encryption_enabled|Bool|
|volume_encryption_key|String|
|workspace_id|String|
|workspace_properties|JSON|