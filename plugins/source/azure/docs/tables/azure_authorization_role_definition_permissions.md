
# Table: azure_authorization_role_definition_permissions
Permission role definition permissions
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|role_definition_cq_id|uuid|Unique CloudQuery ID of azure_authorization_role_definitions table (FK)|
|actions|text[]|Allowed actions|
|not_actions|text[]|Denied actions|
