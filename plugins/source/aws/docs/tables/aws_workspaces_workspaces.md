# Table: aws_workspaces_workspaces

This table shows data for Workspaces Workspaces.

https://docs.aws.amazon.com/workspaces/latest/api/API_Workspace.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|bundle_id|`utf8`|
|computer_name|`utf8`|
|data_replication_settings|`json`|
|directory_id|`utf8`|
|error_code|`utf8`|
|error_message|`utf8`|
|ip_address|`utf8`|
|modification_states|`json`|
|related_workspaces|`json`|
|root_volume_encryption_enabled|`bool`|
|standby_workspaces_properties|`json`|
|state|`utf8`|
|subnet_id|`utf8`|
|user_name|`utf8`|
|user_volume_encryption_enabled|`bool`|
|volume_encryption_key|`utf8`|
|workspace_id|`utf8`|
|workspace_properties|`json`|