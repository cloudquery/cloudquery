
# Table: azure_ad_applications
Application active Directory application information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|app_id|text|The application ID|
|allow_guests_sign_in|boolean|A property on the application to indicate if the application accepts other IDPs or not or partially accepts|
|allow_passthrough_users|boolean|Indicates that the application supports pass through users who have no presence in the resource tenant|
|app_logo_url|text|The url for the application logo image stored in a CDN|
|app_permissions|text[]|The application permissions|
|available_to_other_tenants|boolean|Whether the application is available to other tenants|
|display_name|text|The display name of the application|
|error_url|text|A URL provided by the author of the application to report errors when using the application|
|group_membership_claims|text|Configures the groups claim issued in a user or OAuth 20 access token that the app expects Possible values include: 'None', 'SecurityGroup', 'All'|
|homepage|text|The home page of the application|
|identifier_uris|text[]|A collection of URIs for the application|
|informational_urls_terms_of_service|text|The terms of service URI|
|informational_urls_marketing|text|The marketing URI|
|informational_urls_privacy|text|The privacy policy URI|
|informational_urls_support|text|The support URI|
|is_device_only_auth_supported|boolean|Specifies whether this application supports device authentication without a user The default is false|
|known_client_applications|text[]|Client applications that are tied to this resource application Consent to any of the known client applications will result in implicit consent to the resource application through a combined consent dialog (showing the OAuth permission scopes required by the client and the resource)|
|logout_url|text|the url of the logout page|
|oauth2_allow_implicit_flow|boolean|Whether to allow implicit grant flow for OAuth2|
|oauth2_allow_url_path_matching|boolean|Specifies whether during a token Request Azure AD will allow path matching of the redirect URI against the applications collection of replyURLs The default is false|
|oauth2_require_post_response|boolean|Specifies whether, as part of OAuth 20 token requests, Azure AD will allow POST requests, as opposed to GET requests The default is false, which specifies that only GET requests will be allowed|
|org_restrictions|text[]|A list of tenants allowed to access application|
|optional_claims|jsonb||
|public_client|boolean|Specifies whether this application is a public client (such as an installed application running on a mobile device) Default is false|
|publisher_domain|text|Reliable domain which can be used to identify an application|
|reply_urls|text[]|A collection of reply URLs for the application|
|saml_metadata_url|text|The URL to the SAML metadata for the application|
|sign_in_audience|text|Audience for signing in to the application (AzureADMyOrganization, AzureADAllOrganizations, AzureADAndMicrosoftAccounts)|
|www_homepage|text|The primary Web page|
|additional_properties|jsonb|Unmatched properties from the message are deserialized this collection|
|object_id|text|The object ID|
|deletion_timestamp_time|timestamp without time zone||
|object_type|text|Possible values include: 'ObjectTypeDirectoryObject', 'ObjectTypeApplication', 'ObjectTypeGroup', 'ObjectTypeServicePrincipal', 'ObjectTypeUser'|
