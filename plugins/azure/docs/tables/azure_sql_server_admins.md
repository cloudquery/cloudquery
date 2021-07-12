
# Table: azure_sql_server_admins
ServerAzureADAdministrator azure Active Directory administrator
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_sql_servers table (FK)|
|administrator_type|text|Type of the sever administrator|
|login|text|Login name of the server administrator|
|sid|uuid|SID (object ID) of the server administrator|
|tenant_id|uuid|Tenant ID of the administrator|
|azure_ad_only_authentication|boolean|Azure Active Directory only Authentication enabled|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
