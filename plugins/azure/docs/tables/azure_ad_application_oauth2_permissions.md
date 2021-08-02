
# Table: azure_ad_application_oauth2_permissions
OAuth2Permission represents an OAuth 20 delegated permission scope The specified OAuth 20 delegated permission scopes may be requested by client applications (through the requiredResourceAccess collection on the Application object) when calling a resource application The oauth2Permissions property of the ServicePrincipal entity and of the Application entity is a collection of OAuth2Permission
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|application_cq_id|uuid|Unique ID of azure_ad_applications table (FK)|
|admin_consent_description|text|Permission help text that appears in the admin consent and app assignment experiences|
|admin_consent_display_name|text|Display name for the permission that appears in the admin consent and app assignment experiences|
|id|text|Unique scope permission identifier inside the oauth2Permissions collection|
|is_enabled|boolean|When creating or updating a permission, this property must be set to true (which is the default) To delete a permission, this property must first be set to false At that point, in a subsequent call, the permission may be removed|
|permission_type|text|Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by a Company Administrator Possible values are "User" or "Admin"|
|user_consent_description|text|Permission help text that appears in the end user consent experience|
|user_consent_display_name|text|Display name for the permission that appears in the end user consent experience|
|scope_claim_value|text|The value of the scope claim that the resource application should expect in the OAuth 20 access token|
