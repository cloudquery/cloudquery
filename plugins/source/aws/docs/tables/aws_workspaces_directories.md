# Table: aws_workspaces_directories


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|arn (PK)|String|
|alias|String|
|customer_user_name|String|
|directory_id|String|
|directory_name|String|
|directory_type|String|
|dns_ip_addresses|StringArray|
|iam_role_id|String|
|ip_group_ids|StringArray|
|registration_code|String|
|selfservice_permissions|JSON|
|state|String|
|subnet_ids|StringArray|
|tenancy|String|
|workspace_access_properties|JSON|
|workspace_creation_properties|JSON|
|workspace_security_group_id|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|