
# Table: azure_ad_application_required_resource_accesses
RequiredResourceAccess specifies the set of OAuth 20 permission scopes and app roles under the specified resource that an application requires access to The specified OAuth 20 permission scopes may be requested by client applications (through the requiredResourceAccess collection) when calling a resource application The requiredResourceAccess property of the Application entity is a collection of RequiredResourceAccess
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|application_id|uuid|Unique ID of azure_ad_applications table (FK)|
|additional_properties|jsonb|Unmatched properties from the message are deserialized this collection|
|resource_access|jsonb|The list of OAuth20 permission scopes and app roles that the application requires from the specified resource|
|resource_app_id|text|The unique identifier for the resource that the application requires access to This should be equal to the appId declared on the target resource application|
