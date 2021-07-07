
# Table: azure_ad_application_pre_authorized_applications
PreAuthorizedApplication contains information about pre authorized client application
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|application_cq_id|uuid|Unique ID of azure_ad_applications table (FK)|
|app_id|text|Represents the application id|
|permissions|jsonb|Collection of required app permissions/entitlements from the resource application|
|extensions|jsonb|Collection of extensions from the resource application|
