
# Table: azure_ad_application_app_roles
AppRole
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|application_id|uuid|Unique ID of azure_ad_applications table (FK)|
|resource_id|text|Unique role identifier inside the appRoles collection|
|allowed_member_types|text[]|Specifies whether this app role definition can be assigned to users and groups by setting to 'User', or to other applications (that are accessing this application in daemon service scenarios) by setting to 'Application', or to both|
|description|text|Permission help text that appears in the admin app assignment and consent experiences|
|display_name|text|Display name for the permission that appears in the admin consent and app assignment experiences|
|is_enabled|boolean|When creating or updating a role definition, this must be set to true (which is the default) To delete a role, this must first be set to false At that point, in a subsequent call, this role may be removed|
|role_claim_value|text|Specifies the value of the roles claim that the application should expect in the authentication and access tokens|
