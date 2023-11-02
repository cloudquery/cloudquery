# Table: aws_workspaces_directories

This table shows data for Workspaces Directories.

https://docs.aws.amazon.com/workspaces/latest/api/API_WorkspaceDirectory.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|alias|`utf8`|
|certificate_based_auth_properties|`json`|
|customer_user_name|`utf8`|
|directory_id|`utf8`|
|directory_name|`utf8`|
|directory_type|`utf8`|
|dns_ip_addresses|`list<item: utf8, nullable>`|
|iam_role_id|`utf8`|
|ip_group_ids|`list<item: utf8, nullable>`|
|registration_code|`utf8`|
|saml_properties|`json`|
|selfservice_permissions|`json`|
|state|`utf8`|
|subnet_ids|`list<item: utf8, nullable>`|
|tenancy|`utf8`|
|workspace_access_properties|`json`|
|workspace_creation_properties|`json`|
|workspace_security_group_id|`utf8`|