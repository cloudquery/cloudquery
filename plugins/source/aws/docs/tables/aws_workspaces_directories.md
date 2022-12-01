# Table: aws_workspaces_directories

https://docs.aws.amazon.com/workspaces/latest/api/API_WorkspaceDirectory.html

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
|alias|String|
|certificate_based_auth_properties|JSON|
|customer_user_name|String|
|directory_id|String|
|directory_name|String|
|directory_type|String|
|dns_ip_addresses|StringArray|
|iam_role_id|String|
|ip_group_ids|StringArray|
|registration_code|String|
|saml_properties|JSON|
|selfservice_permissions|JSON|
|state|String|
|subnet_ids|StringArray|
|tenancy|String|
|workspace_access_properties|JSON|
|workspace_creation_properties|JSON|
|workspace_security_group_id|String|