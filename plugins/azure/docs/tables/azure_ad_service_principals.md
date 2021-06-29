
# Table: azure_ad_service_principals
ServicePrincipal active Directory service principal information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|account_enabled|boolean|whether or not the service principal account is enabled|
|alternative_names|text[]|alternative names|
|app_display_name|text|The display name exposed by the associated application|
|app_id|text|The application ID|
|app_owner_tenant_id|text|Application owner id|
|app_role_assignment_required|boolean|Specifies whether an AppRoleAssignment to a user or group is required before Azure AD will issue a user or access token to the application|
|display_name|text|The display name of the service principal|
|error_url|text|A URL provided by the author of the associated application to report errors when using the application|
|homepage|text|The URL to the homepage of the associated application|
|logout_url|text|A URL provided by the author of the associated application to logout|
|preferred_token_signing_key_thumbprint|text|The thumbprint of preferred certificate to sign the token|
|publisher_name|text|The publisher's name of the associated application|
|reply_urls|text[]|The URLs that user tokens are sent to for sign in with the associated application  The redirect URIs that the oAuth 20 authorization code and access tokens are sent to for the associated application|
|saml_metadata_url|text|The URL to the SAML metadata of the associated application|
|service_principal_names|text[]|A collection of service principal names|
|service_principal_type|text|the type of the service principal|
|tags|text[]|Optional list of tags that you can apply to your service principals Not nullable|
|additional_properties|jsonb|Unmatched properties from the message are deserialized this collection|
|object_id|text|The object ID|
|deletion_timestamp_time|timestamp without time zone|The time at which the directory object was deleted.|
|object_type|text|Possible values include: 'ObjectTypeDirectoryObject', 'ObjectTypeApplication', 'ObjectTypeGroup', 'ObjectTypeServicePrincipal', 'ObjectTypeUser'|
